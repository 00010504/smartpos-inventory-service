package models

type OrderItemES struct {
	Id              string   `json:"id"`
	ProductId       string   `json:"product_id"`
	ProductName     string   `json:"product_name"`
	Barcode         []string `json:"barcode"`
	ExpectedAmount  int32    `json:"expected_amount"`
	Cost            int32    `json:"cost"`
	Discount        int32    `json:"discount"`
	SoldAmount      int32    `json:"sold_amount"`
	CreatedAt       string   `json:"created_at"`
	ArrivedAmount   int32    `json:"arrived_amount"`
	SupplierOrderId string   `json:"supplier_order_id"`
}
