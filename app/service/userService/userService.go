package userService

import (
	"errors"
	"ransmart_pay/app/helper/helper"
	"ransmart_pay/app/models/payModel"
	"ransmart_pay/app/models/userModel"
	"ransmart_pay/app/repository"

	"gorm.io/gorm"
)

type service struct {
	repository repository.Repository
	db         *gorm.DB
}

func NewService(repository repository.Repository, db *gorm.DB) *service {
	return &service{repository, db}
}

func (s *service) FindByUsername(username string) (userModel.User, error) {
	data, err := s.repository.IUserRepository.FindByUsername(username)
	if err != nil {
		return data, errors.New("user tidak ditemukan")
	}
	return data, nil
}

func (s *service) Create(user userModel.User) (err error) {
	newPassword := helper.Encode([]byte(user.Password))
	user.Password = string(newPassword)

	tx := s.db.Begin()
	err = s.repository.IUserRepository.Create(tx, user)
	if err != nil {
		return errors.New("gagal menambah data user")
	}

	userPay := payModel.PayModel{
		Username: user.Username,
		Saldo:    0,
	}
	err = s.repository.IPayRepository.Create(tx, userPay)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func (s *service) Update(id int, User userModel.User) (err error) {
	_, err = s.repository.IUserRepository.Update(id, User)
	if err != nil {
		return errors.New("gagal mengupdate user")
	}
	return
}
