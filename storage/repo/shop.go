package repo

import (
	"genproto/common"
	"genproto/corporate_service"
)

type ShopI interface {
	Upsert(*common.ShopCreatedModel) error
	GetById(entity *common.RequestID) (*corporate_service.Shop, error)
}
