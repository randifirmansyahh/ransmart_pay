package service

import (
	"ransmart_pay/app/service/payHistoryService"
	"ransmart_pay/app/service/payService"
	"ransmart_pay/app/service/userService"
)

type Service struct {
	IUserService       userService.IUserService
	IPayService        payService.IPayService
	IPayHistoryService payHistoryService.IPayHistoryService
}
