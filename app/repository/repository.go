package repository

import (
	"ransmart_pay/app/repository/payHistoryRepository"
	"ransmart_pay/app/repository/payRepository"
	"ransmart_pay/app/repository/userRepository"
)

type Repository struct {
	IUserRepository       userRepository.IUserRepository
	IPayRepository        payRepository.IPayRepository
	IPayHistoryRepository payHistoryRepository.IPayHistoryRepo
}
