package payHistoryRepository

import (
	"ransmart_pay/app/models/payHistoryModel"

	"gorm.io/gorm"
)

type IPayHistoryRepo interface {
	FindByUsername(username string) (data []payHistoryModel.PayHistoryModel, err error)
	Create(tx *gorm.DB, data payHistoryModel.PayHistoryModel) (err error)
	FindByOrderId(orderId int) (data payHistoryModel.PayHistoryModel, err error)
}
