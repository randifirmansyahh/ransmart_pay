package payHistoryService

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"ransmart_pay/app/helper/helper"
	"ransmart_pay/app/helper/tokenHelper"
	"ransmart_pay/app/httpRequest"
	"ransmart_pay/app/models/orderModel"
	"ransmart_pay/app/models/payHistoryModel"
	"ransmart_pay/app/repository"
	"strconv"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type service struct {
	repository repository.Repository
	db         *gorm.DB
}

func NewService(repository repository.Repository, db *gorm.DB) *service {
	return &service{repository, db}
}

func (s *service) FindByUsername(username string) (data []payHistoryModel.PayHistoryModel, err error) {
	_, err = s.repository.IUserRepository.FindByUsername(username)
	if err != nil {
		return nil, errors.New("username tersebut tidak ditemukan")
	}

	data, err = s.repository.IPayHistoryRepository.FindByUsername(username)
	if err != nil {
		return nil, errors.New("data dengan username tersebut tidak ditemukan")
	}
	return
}

func (s *service) Create(data payHistoryModel.PayHistoryReq) (err error) {
	// find order from ransmart-product
	// generate JWT
	ISS := os.Getenv("JWT_ISS")
	AUD := os.Getenv("JWT_AUD")
	JWT_SECRET_KEY := os.Getenv("JWT_SECRET_KEY")
	JWT_EXPIRATION_DURATION_DAY := os.Getenv("JWT_EXPIRATION_DURATION_DAY")
	newWaktu, _ := strconv.Atoi(JWT_EXPIRATION_DURATION_DAY)
	expiredTime := helper.ExpiredTime(newWaktu)

	jwt, err := tokenHelper.BuatJWT(ISS, AUD, JWT_SECRET_KEY, expiredTime)
	if err != nil {
		log.Error().Msgf("error buat jwt: %v", err)
		return errors.New("gagal generate token")
	}

	// set header
	header := map[string]string{
		"Authorization": "Bearer " + jwt,
	}

	_, err = s.repository.IPayHistoryRepository.FindByOrderId(data.OrderId)
	if err == nil {
		log.Error().Msgf("orderId %v sudah ada", data.OrderId)
		return errors.New("order telah dibayar sebelumnya")
	}

	_, err = s.repository.IUserRepository.FindByUsername(data.Username)
	if err != nil {
		log.Error().Msgf("error get user: %v", err)
		return errors.New("username tidak ditemukan")
	}

	findSaldo, err := s.repository.IPayRepository.FindByUsername(data.Username)
	if err != nil {
		log.Error().Msgf("error get saldo: %v", err)
		return errors.New("gagal mengambil data saldo")
	}

	// find order from ransmart_product
	log.Printf("find order di ransmart-product")
	url := "https://ransmart-product.herokuapp.com/order/id/" + fmt.Sprint(data.OrderId)
	code, result, err := httpRequest.HTTPResponse("GET", url, "", header)
	if err != nil || code != 200 {
		log.Error().Msgf("error get order: %v", err)
		return errors.New("data order tidak ditemukan")
	}

	// fill result to data
	log.Printf("berhasil get order")
	var resOrder orderModel.OrderHttpResponse
	json.Unmarshal([]byte(result), &resOrder)

	TOTAL := resOrder.Data.Total
	if TOTAL > findSaldo.Saldo {
		log.Error().Msgf("saldo tidak cukup")
		return errors.New("saldo tidak mencukupi, saldo = Rp." + fmt.Sprint(findSaldo.Saldo) + ", pembayaran = Rp." + fmt.Sprint(TOTAL))
	}

	request := payHistoryModel.PayHistoryModel{
		Username:    data.Username,
		OrderId:     data.OrderId,
		Pengeluaran: TOTAL,
	}

	tx := s.db.Begin()
	err = s.repository.IPayHistoryRepository.Create(tx, request)
	if err != nil {
		tx.Rollback()
		log.Error().Msgf("error create payHistory: %v", err)
		return errors.New("gagal menambahkan data order")
	}

	findSaldo.Saldo -= TOTAL
	err = s.repository.IPayRepository.UpdateSaldo(tx, findSaldo)
	if err != nil {
		tx.Rollback()
		log.Error().Msgf("error update saldo: %v", err)
		return errors.New("gagal membayar order")
	}

	urlUpdate := "https://ransmart-product.herokuapp.com/order/" + fmt.Sprint(findSaldo.Saldo) + "/" + fmt.Sprint(data.OrderId)
	code, result, err = httpRequest.HTTPResponse("PUT", urlUpdate, "", header)
	if err != nil || code != 200 {
		tx.Rollback()
		log.Error().Msgf("error get order: %v, result : ", err, result)
		return errors.New("gagal mengupdate data order")
	}

	tx.Commit()
	return
}

func (s *service) FindByOrderId(orderId int) (data payHistoryModel.PayHistoryModel, err error) {
	data, err = s.repository.IPayHistoryRepository.FindByOrderId(orderId)
	if err != nil {
		return data, errors.New("data dengan orderId tersebut tidak ditemukan")
	}
	return
}
