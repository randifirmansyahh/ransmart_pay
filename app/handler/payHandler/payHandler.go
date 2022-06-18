package payHandler

import (
	"encoding/json"
	"net/http"
	"ransmart_pay/app/helper/response"
	"ransmart_pay/app/models/payModel"
	"ransmart_pay/app/service"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
)

var (
	HandlerName = "ransmart_pay"
)

type payHandler struct {
	service service.Service
}

func NewPayHandler(payService service.Service) *payHandler {
	return &payHandler{payService}
}

func (h *payHandler) FindByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	findData, err := h.service.IPayService.FindByUsername(username)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response.Response(w, http.StatusOK, response.MsgGetDetail(true, HandlerName), findData)
}

func (h *payHandler) UpdateSaldo(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var datarequest payModel.PayRequest
	if err := decoder.Decode(&datarequest); err != nil {
		response.Response(w, http.StatusBadRequest, "Data harus berupa json / request kurang lengkap", nil)
		return
	}

	err := h.checkValidation(datarequest)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// find user
	findUser, err := h.service.IUserService.FindByUsername(datarequest.Username)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// find saldo
	findSaldo, err := h.service.IPayService.FindByUsername(findUser.Username)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	findSaldo.Saldo += datarequest.Saldo

	// update
	err = h.service.IPayService.UpdateSaldo(findSaldo)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// response success
	response.Response(w, http.StatusOK, "terima kasih telah mengisi saldo", nil)
}

func (h *payHandler) checkValidation(datarequest payModel.PayRequest) (err error) {
	validate := validator.New()
	if err = validate.Struct(datarequest); err != nil {
		if errors := err.(validator.ValidationErrors); errors != nil {
			return errors
		}
	}
	return
}
