package orderModel

import (
	"ransmart_pay/app/models/modelHelper"
	"ransmart_pay/app/models/productModel"
)

type Order struct {
	Id          int                  `gorm:"primaryKey;autoIncrement;" json:"id"`
	Username    string               `json:"username"`
	Product_Id  int                  `json:"product_id"`
	Product     productModel.Product `json:"product" gorm:"foreignKey:Product_Id;references:id"`
	Qty         int                  `json:"qty"`
	Total       int                  `json:"total"`
	OrderStatus bool                 `json:"order_status"`
	modelHelper.DateAuditModel
}

type OrderReq struct {
	Username  string `json:"username" validate:"required"`
	ProductId int    `json:"product_id" validate:"required"`
	Qty       int    `json:"qty" validate:"required,number"`
}

type Header struct {
	Status       string `json:"status"`
	ResponseCode string `json:"response_code"`
	Message      string `json:"message"`
}

type OrderHttpResponse struct {
	Header
	Data OrderResponse `json:"data"`
}

type OrderResponse struct {
	Id          int                  `json:"id"`
	Username    string               `json:"username"`
	Product     productModel.Product `json:"product"`
	Qty         int                  `json:"qty"`
	Total       int                  `json:"total"`
	OrderStatus bool                 `json:"order_status"`
	modelHelper.DateAuditModel
}
