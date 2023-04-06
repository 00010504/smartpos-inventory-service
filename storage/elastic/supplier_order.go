package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"genproto/common"
	"genproto/inventory_service"
	"io/ioutil"

	"github.com/Invan2/invan_inventory_service/config"
	"github.com/Invan2/invan_inventory_service/models"
	"github.com/Invan2/invan_inventory_service/pkg/logger"
	"github.com/Invan2/invan_inventory_service/storage/repo"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/golang/protobuf/jsonpb"
	"github.com/pkg/errors"
)

type supplierOrderItemRepo struct {
	db  *elasticsearch.Client
	log logger.Logger
}

func NewSupplierOrderRepo(log logger.Logger, db *elasticsearch.Client) repo.SupplierOrderItemESI {
	return &supplierOrderItemRepo{
		db:  db,
		log: log,
	}
}

type H map[string]interface{}

func (s *supplierOrderItemRepo) CreateMultiple(req *inventory_service.CreateSupplierOrderElasticRequest) error {
	var (
		buf bytes.Buffer
	)

	if len(req.GetItemsEs()) == 0 {
		return errors.New("no supplierOrderItems were provided")
	}

	indexName := config.ElasticOrderItemIndex + req.Request.CompanyId

	if !exists(s.db, indexName) {
		res, err := s.db.Indices.Create(indexName)
		if err != nil {
			return errors.Wrap(err, "error while create SupplierOrderItems index on elastic")
		}

		if err := checkResponseCodeToSuccess(res.StatusCode); err != nil {
			return err
		}
	}

	for _, item := range req.GetItemsEs() {

		meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%s" }%s`, item.Id, "\n"))

		m := jsonpb.Marshaler{
			EmitDefaults: true,
			OrigName:     true,
		}

		var bufItem bytes.Buffer

		if err := m.Marshal(&bufItem, item); err != nil {
			return errors.Wrap(err, "error while marshaling supplier order item")
		}

		data := bufItem.Bytes()

		data = append(data, "\n"...)

		buf.Grow(len(meta) + len(data))
		buf.Write(meta)
		buf.Write(data)

	}

	res, err := s.db.Bulk(bytes.NewReader(buf.Bytes()), s.db.Bulk.WithIndex(indexName), s.db.Bulk.WithRefresh("wait_for"))
	if err != nil {
		return errors.Wrap(err, "Failed to bulk insert supplier order item")
	}
	defer res.Body.Close()

	if res.IsError() {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		s.log.Error("errror while create ", logger.Any("res", string(data)))
		return errors.New("error while create supplier_order_item on elastic")
	}

	if err := checkResponseCodeToSuccess(res.StatusCode); err != nil {
		return err
	}

	return nil
}

func (s *supplierOrderItemRepo) UpsertOrderItem(req *inventory_service.SupplierOrderItemES) error {
	var (
		body bytes.Buffer
	)
	index := config.ElasticOrderItemIndex + req.CompanyId

	if !exists(s.db, index) {
		res, err := s.db.Indices.Create(index)
		if err != nil {
			return errors.Wrap(err, "error while create index")
		}

		if err := checkResponseCodeToSuccess(res.StatusCode); err != nil {
			return err
		}
	}

	err := config.JSONPBMarshaler.Marshal(&body, &inventory_service.UpsertOrderItemES{
		Doc: &inventory_service.SupplierOrderItemES{
			Id:              req.Id,
			Barcode:         req.Barcode,
			CompanyId:       req.CompanyId,
			ProductId:       req.ProductId,
			ExpectedAmount:  req.ExpectedAmount,
			ArrivedAmount:   req.ArrivedAmount,
			Cost:            req.Cost,
			Discount:        req.Discount,
			SoldAmount:      req.SoldAmount,
			CreatedAt:       req.CreatedAt,
			ProductName:     req.ProductName,
			SupplierOrderId: req.SupplierOrderId,
		},
		DocAsUpsert: true,
	})
	if err != nil {
		return errors.Wrap(err, "error while marshaling, jsonpb")
	}

	meta := []byte(fmt.Sprintf(`{ "update": { "_index": "%s", "_id" : "%s", "retry_on_conflict": 1 } }%s`, index, req.Id, "\n"))

	body.Grow(len("\n"))
	body.Write(bytes.NewBufferString("\n").Bytes())

	body.Grow(len(meta) + len(body.Bytes()))
	body.Write(meta)
	body.Write(body.Bytes())
	res, err := s.db.Update(
		index,
		req.Id,
		bytes.NewReader(body.Bytes()),
		s.db.Update.WithRefresh("wait_for"),
		s.db.Update.WithPretty(),
	)
	if err != nil {
		return errors.Wrap(err, "Failed to bulk insert orderItem")
	}
	defer res.Body.Close()

	if res.IsError() {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		s.log.Error("errror while create many ", logger.Any("res", string(data)))
		return errors.New("error while create ortedItem on elastic")
	}

	if err := checkResponseCodeToSuccess(res.StatusCode); err != nil {
		return err
	}

	return nil
}

func (s *supplierOrderItemRepo) DeleteSupplierOrderItemById(req *common.RequestID) (*common.Empty, error) {
	index := config.ElasticOrderItemIndex + req.Request.CompanyId

	var (
		buf bytes.Buffer
	)

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"ids": map[string]interface{}{
				"values": req.Id,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, errors.Wrap(err, "Error encoding query")
	}

	res, err := s.db.DeleteByQuery(
		[]string{index},
		&buf,
		s.db.DeleteByQuery.WithRefresh(true),
	)
	if err != nil {
		return nil, err
	}

	if err := checkResponseCodeToSuccess(res.StatusCode); err != nil {
		return nil, err
	}

	return &common.Empty{}, nil
}

func makeGettAllSearchRequest(req *inventory_service.GetAllSupplierOrderItemsRequest) H {

	must := make([]H, 0)
	bool := make(H)
	query := make(H)

	must = append(must, H{
		"term": H{
			"supplier_order_id.keyword": H{
				"value": req.SupplierOrderId,
			},
		},
	})

	if req.Search != "" {
		must = append(must, H{
			"query_string": H{
				"query":            getSearchString(req.Search),
				"default_operator": "AND",
				"fields":           []string{"barcode", "product_name"},
			},
		})
	}

	if len(must) > 0 {
		bool["must"] = must
	}

	if len(bool) > 0 {
		query["bool"] = bool
	}

	sort := make([]H, 0)

	if req.SortBy == "" {
		sort = append(sort, H{
			"created_at.keyword": H{
				"order": "desc",
			},
		})
	}

	return H{
		"query": query,
		"sort":  sort,
	}

}

func (s *supplierOrderItemRepo) GetAllOrderItems(entity *inventory_service.GetAllSupplierOrderItemsRequest) (*inventory_service.GetAllSupplierOrderItemsResponse, error) {

	var (
		res = inventory_service.GetAllSupplierOrderItemsResponse{
			Data:       make([]*inventory_service.CreateSupplierOrderItemElasticResponse, 0),
			TotalCount: 0,
		}

		r    map[string]interface{}
		buf  bytes.Buffer
		size = int(entity.Limit)
		from = int((entity.Page - 1) * entity.Limit)
	)

	req := makeGettAllSearchRequest(entity)

	if err := json.NewEncoder(&buf).Encode(req); err != nil {
		return nil, err
	}

	fmt.Println(buf.String())

	response, err := s.db.Search(
		s.db.Search.WithContext(context.Background()),
		s.db.Search.WithIndex(config.ElasticOrderItemIndex+entity.Request.CompanyId),
		s.db.Search.WithBody(&buf),
		s.db.Search.WithTrackTotalHits(true),
		s.db.Search.WithFrom(from),
		s.db.Search.WithSize(size),
	)
	if err != nil {
		return nil, errors.Wrap(err, "error while get documents on elastic")
	}

	if response.IsError() {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		s.log.Error("errror while get all products ", logger.Any("res", string(data)))
		return nil, errors.New("error while get  products on elastic " + string(data))
	}
	if response.StatusCode != 200 {
		return nil, errors.New("error while get documents on elastic. code: " + response.Status())
	}

	err = json.NewDecoder(response.Body).Decode(&r)
	if err != nil {
		return nil, errors.Wrap(err, "error while json.decode elastic res.Body")
	}

	for _, source := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {

		product := models.OrderItemES{}

		jsonString, _ := json.Marshal(source.(map[string]interface{})["_source"])

		err = json.Unmarshal(jsonString, &product)
		if err != nil {
			return nil, errors.Wrap(err, "error while json.Unmarshal jsonString &import")
		}

		res.Data = append(res.Data, &inventory_service.CreateSupplierOrderItemElasticResponse{
			Id:              product.Id,
			ProductId:       product.ProductId,
			ProductName:     product.ProductName,
			Barcode:         product.Barcode,
			ExpectedAmount:  float32(product.ExpectedAmount),
			ArrivedAmount:   float32(product.ArrivedAmount),
			Cost:            float32(product.Cost),
			Discount:        float32(product.Discount),
			SoldAmount:      float32(product.SoldAmount),
			CreatedAt:       product.CreatedAt,
			SupplierOrderId: product.SupplierOrderId,
		})
	}

	res.TotalCount = int32(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

	return &res, nil
}
