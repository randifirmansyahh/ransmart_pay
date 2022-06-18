package payService

import "ransmart_pay/app/models/payModel"

type IPayService interface {
	FindByUsername(username string) (data payModel.PayModel, err error)
	Create(data payModel.PayModel) (err error)
	UpdateSaldo(data payModel.PayModel) (err error)
}
