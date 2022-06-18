package payHistoryHandler

import (
	"encoding/json"
	"net/http"
	"ransmart_pay/app/helper/requestHelper"
	"ransmart_pay/app/helper/response"
	"ransmart_pay/app/models/payHistoryModel"
	"ransmart_pay/app/service"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
)

var (
	HandlerName = "pay history"
)

type payHistoryHandler struct {
	service service.Service
}

func NewPayHistoryHandler(payHistoryService service.Service) *payHistoryHandler {
	return &payHistoryHandler{payHistoryService}
}

func (h *payHistoryHandler) FindByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	findData, err := h.service.IPayHistoryService.FindByUsername(username)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response.Response(w, http.StatusOK, response.MsgGetDetail(true, HandlerName), findData)
}

func (h *payHistoryHandler) FindByOrderId(w http.ResponseWriter, r *http.Request) {
	orderId := chi.URLParam(r, "order_id")

	newId, err := requestHelper.CheckIDInt(orderId)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	data, err := h.service.IPayHistoryService.FindByOrderId(newId)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// response success
	response.Response(w, http.StatusOK, response.MsgGetDetail(true, HandlerName), data)
}

func (h *payHistoryHandler) PayOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var datarequest payHistoryModel.PayHistoryReq
	if err := decoder.Decode(&datarequest); err != nil {
		response.Response(w, http.StatusBadRequest, "Data harus berupa json / request kurang lengkap", nil)
		return
	}

	err := h.checkValidation(datarequest)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = h.service.IPayHistoryService.Create(datarequest)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// response success
	response.Response(w, http.StatusOK, "terima kasih telah membayar order tersebut", nil)
}

func (h *payHistoryHandler) checkValidation(datarequest payHistoryModel.PayHistoryReq) (err error) {
	validate := validator.New()
	if err = validate.Struct(datarequest); err != nil {
		if errors := err.(validator.ValidationErrors); errors != nil {
			return errors
		}
	}
	return
}
