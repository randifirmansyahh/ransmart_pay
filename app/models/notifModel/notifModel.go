package notifModel

type NotifHttpResponse struct {
	OrderId int           `json:"order_id"`
	Message string        `json:"message"`
	Data    NotifResponse `json:"data"`
}

type NotifResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	ProductId int    `json:"product_id"`
	Qty       int    `json:"qty"`
	Total     int    `json:"total"`
}
