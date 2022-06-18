package payHistoryRepository

import (
	"ransmart_pay/app/models/payHistoryModel"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IPayHistoryRepo {
	return &repository{db: db}
}

func (r *repository) FindByUsername(username string) (data []payHistoryModel.PayHistoryModel, err error) {
	err = r.db.Find(&data, "username = ?", username).Error
	return data, err
}

func (r *repository) Create(tx *gorm.DB, data payHistoryModel.PayHistoryModel) (err error) {
	if tx != nil {
		err = tx.Create(&data).Error
		return
	}
	err = r.db.Create(&data).Error
	return err
}

func (r *repository) FindByOrderId(orderId int) (data payHistoryModel.PayHistoryModel, err error) {
	err = r.db.First(&data, "order_id = ?", orderId).Error
	return data, err
}
