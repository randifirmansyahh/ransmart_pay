package payHistoryModel

type (
	PayHistoryModel struct {
		Id          int    `json:"id"`
		Username    string `json:"username"`
		OrderId     int    `json:"order_id"`
		Pengeluaran int    `json:"pengeluaran"`
	}

	PayHistoryModelResponse struct {
		PayHistoryModel  []PayHistoryModel `json:"pay_history"`
		TotalPengeluaran int               `json:"total_pengeluaran"`
		Saldo            int               `json:"saldo"`
	}

	PayHistoryReq struct {
		Username string `json:"username" validate:"required"`
		OrderId  int    `json:"order_id" validate:"required"`
	}
)

func (PayHistoryModel) TableName() string {
	return "pay_history"
}
