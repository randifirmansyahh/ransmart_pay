package payRepository

import (
	"ransmart_pay/app/models/payModel"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IPayRepository {
	return &repository{db: db}
}

func (r *repository) FindByUsername(username string) (data payModel.PayModel, err error) {
	err = r.db.First(&data, "username = ?", username).Error
	return
}

func (r *repository) Create(tx *gorm.DB, data payModel.PayModel) (err error) {
	if tx != nil {
		err = tx.Create(&data).Error
		return
	}
	err = r.db.Create(&data).Error
	return
}

func (r *repository) UpdateSaldo(tx *gorm.DB, data payModel.PayModel) (err error) {
	if tx != nil {
		err = tx.Save(&data).Error
		return
	}
	err = r.db.Save(&data).Error
	return
}
