package payHistoryService

import "ransmart_pay/app/models/payHistoryModel"

type IPayHistoryService interface {
	FindByUsername(username string) (data []payHistoryModel.PayHistoryModel, err error)
	Create(data payHistoryModel.PayHistoryReq) (err error)
	FindByOrderId(orderId int) (data payHistoryModel.PayHistoryModel, err error)
}
