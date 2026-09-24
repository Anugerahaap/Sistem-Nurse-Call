package device

import (
	"backend/model/web"
	"backend/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	service   Service
	validator *validator.Validate
}

func NewHandler(service Service, validator *validator.Validate) *Handler {
	return &Handler{service: service, validator: validator}
}

func (h *Handler) RegisterRoute(router *httprouter.Router) {
	router.POST("/api/devices", h.handleRegisterNewDevice)
	router.GET("/api/devices", h.handleGetDevices)
	router.GET("/api/devices/:device-id", h.handleGetDeviceByID)
	router.PATCH("/api/devices", h.handleUpdate)
}

func (h *Handler) handleRegisterNewDevice(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var payload *web.DevicePayload

	if err := utils.ParseJson(r, &payload); err != nil {

		utils.JsonInternalError(w, err.Error(), nil)
		return
	}

	if err := utils.ValidateStruct(h.validator, payload); len(err) > 0 {
		utils.WriteJson(w, http.StatusBadRequest, "bad request", "Required parameter is missing", err)
		return
	}
	if err := h.service.RegisterNewDevice(r.Context(), payload); err != nil {
		switch err {
		case utils.DevicesIdAlrRegistered:
			utils.JsonConflict(w, fmt.Sprintf("Device :%s already registered,try another one", payload.DeviceId), nil)
			return
		default:
			utils.JsonInternalError(w, err.Error(), nil)
			return
		}
	}

	utils.WriteJson(w, http.StatusOK, "status ok", "Register device succesful", &payload)
}

func (h *Handler) handleGetDevices(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	devices, err := h.service.GetDevices(r.Context())
	if err != nil {
		switch err {
		case utils.NotFoundDevices:
			utils.JsonNotFound(w, "No devices found ,try creating one!", nil)
			return
		default:
			utils.JsonInternalError(w, err.Error(), nil)
			return
		}

	}

	utils.WriteJson(w, http.StatusOK, "status ok", "", devices)

}

func (h *Handler) handleGetDeviceByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	deviceID := params.ByName("device-id")
	device, err := h.service.GetDeviceByID(r.Context(), deviceID)
	if err != nil {
		switch err {
		case utils.NotFoundDevices:
			utils.JsonNotFound(w, "No devices found ,try creating one!", nil)
			return
		default:
			utils.JsonInternalError(w, err.Error(), nil)
			return

		}
	}

	utils.WriteJson(w, 200, "status ok", "", device)
}

func (h *Handler) handleUpdate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	query := r.URL.Query().Get("action")

	switch query {
	case "update-room":
		var payload *web.DevicePayload

		if err := utils.ParseJson(r, &payload); err != nil {

			utils.JsonInternalError(w, err.Error(), nil)
			return
		}
		if err := utils.ValidateStruct(h.validator, payload); len(err) > 0 {
			utils.WriteJson(w, http.StatusBadRequest, "bad request", "Required parameter is missing", err)
			return
		}
		if err := h.service.UpdateRoomName(r.Context(), payload); err != nil {
			switch err {
			case utils.DeviceSameParameter:
				utils.JsonConflict(w, err.Error(), nil)
				return
			case utils.DevicesIdAlrRegistered:
				utils.JsonConflict(w, fmt.Sprintf("device-id:%s sudah terdaftar,tidak dapat memiliki nama device yang sama", payload.DeviceId), nil)
				return
			case utils.NotFoundDevices:
				utils.JsonNotFound(w, "No devices found ,make sure device-id is correct", nil)
				return
			default:
				utils.JsonInternalError(w, err.Error(), nil)
				return

			}
		}

		//get the latest updated version of devices
		device, err := h.service.GetDeviceByID(r.Context(), payload.DeviceId)
		if err != nil {

			switch err {
			case utils.NotFoundDevices:
				utils.JsonNotFound(w, "No devices found ,try creating one!", nil)
				return
			default:
				utils.JsonInternalError(w, err.Error(), nil)
				return

			}
		}

		utils.WriteJson(w, 200, "status ok", "", device)
	case "update-id":
		value := r.URL.Query().Get("value")

		var payload *web.DevicePayload

		if err := utils.ParseJson(r, &payload); err != nil {

			utils.JsonInternalError(w, err.Error(), nil)
			return
		}
		if err := utils.ValidateStruct(h.validator, payload); len(err) > 0 {
			utils.WriteJson(w, http.StatusBadRequest, "bad request", "Required parameter is missing", err)
			return
		}

		if err := h.service.UpdateDeviceID(r.Context(), value, payload); err != nil {
			switch err {
			case utils.DeviceSameParameter:
				utils.JsonConflict(w, err.Error(), nil)
				return
			case utils.DevicesIdAlrRegistered:
				utils.JsonConflict(w, fmt.Sprintf("device-id:%s sudah terdaftar,tidak dapat memiliki nama device yang sama", payload.DeviceId), nil)
				return
			case utils.NotFoundDevices:
				utils.JsonNotFound(w, "No devices found ,make sure device-id is correct", nil)
				return
			default:
				utils.JsonInternalError(w, err.Error(), nil)
				return

			}
		}
		//get the latest updated version of devices
		device, err := h.service.GetDeviceByID(r.Context(), payload.DeviceId)
		if err != nil {

			switch err {
			case utils.NotFoundDevices:
				utils.JsonNotFound(w, "No devices found ,try creating one!", nil)
				return
			default:
				utils.JsonInternalError(w, err.Error(), nil)
				return

			}
		}

		utils.WriteJson(w, 200, "status ok", "", device)
	default:
		utils.JsonConflict(w, "query parameter not allowed", nil)
	}

}
