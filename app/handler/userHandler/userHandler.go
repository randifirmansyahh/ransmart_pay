package userHandler

import (
	"encoding/json"
	"net/http"
	"ransmart_pay/app/helper/requestHelper"
	"ransmart_pay/app/helper/response"
	"ransmart_pay/app/models/userModel"
	"ransmart_pay/app/service"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

var (
	HandlerName = "User"
	paramName   = "id"
)

type userHandler struct {
	service service.Service
}

func NewUserHandler(userService service.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	// decode and fill to model
	decoder := json.NewDecoder(r.Body)
	var datarequest userModel.User
	if err := decoder.Decode(&datarequest); err != nil {
		response.Response(w, http.StatusBadRequest, "Data harus berupa json / request kurang lengkap", nil)
		return
	}

	validate := validator.New()
	err := validate.Struct(datarequest)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if errors != nil {
			response.Response(w, http.StatusBadRequest, errors.Error(), nil)
			return
		}
	}

	// insert
	err = h.service.IUserService.Create(datarequest)
	if err != nil {
		response.Response(w, http.StatusBadRequest, "Username sudah digunakan", nil)
		return
	}

	// response success
	response.Response(w, http.StatusOK, response.MsgCreate(true, HandlerName), nil)
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, paramName)

	// check id
	newId, err := requestHelper.CheckIDInt(id)
	if err != nil {
		response.Response(w, http.StatusBadRequest, "ID harus berupa angka", nil)
	}

	// decode and fill to model
	decoder := json.NewDecoder(r.Body)
	var datarequest userModel.User
	if err := decoder.Decode(&datarequest); err != nil {
		response.Response(w, http.StatusBadRequest, "Data harus berupa json / request kurang lengkap", nil)
		return
	}

	// update
	err = h.service.IUserService.Update(newId, datarequest)
	if err != nil {
		response.Response(w, http.StatusNotFound, "Data dengan ID tersebut tidak ditemukan", nil)
		return
	}

	// response success
	response.Response(w, http.StatusOK, response.MsgUpdate(true, HandlerName), nil)
}
