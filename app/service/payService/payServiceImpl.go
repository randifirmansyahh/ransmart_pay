package payService

import (
	"errors"
	"ransmart_pay/app/models/payModel"
	"ransmart_pay/app/repository"
)

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) FindByUsername(username string) (data payModel.PayModel, err error) {
	data, err = s.repository.IPayRepository.FindByUsername(username)
	if err != nil {
		return data, errors.New("data dengan username tersebut tidak ditemukan")
	}
	return
}

func (s *service) Create(data payModel.PayModel) (err error) {
	err = s.repository.IPayRepository.Create(nil, data)
	if err != nil {
		return errors.New("gagal menambah data pay")
	}
	return err
}

func (s *service) UpdateSaldo(data payModel.PayModel) (err error) {
	err = s.repository.IPayRepository.UpdateSaldo(nil, data)
	if err != nil {
		return errors.New("data dengan username tersebut tidak ditemukan")
	}
	return
}
