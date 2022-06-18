package payModel

type (
	PayModel struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Saldo    int    `json:"saldo"`
	}

	PayRequest struct {
		Username string `json:"username" validate:"required"`
		Saldo    int    `json:"saldo" validate:"required,numeric"`
	}

	PayResponse struct {
		Saldo int `json:"saldo"`
	}
)

func (PayModel) TableName() string {
	return "ransmart_pay"
}
