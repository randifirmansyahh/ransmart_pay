package payRepository

import (
	"ransmart_pay/app/models/payModel"

	"gorm.io/gorm"
)

type IPayRepository interface {
	FindByUsername(username string) (data payModel.PayModel, err error)
	Create(tx *gorm.DB, data payModel.PayModel) (err error)
	UpdateSaldo(tx *gorm.DB, data payModel.PayModel) (err error)
}
