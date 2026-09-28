package device

import (
	"backend/middleware"
	"backend/model/web"
	"backend/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	service    Service
	validator  *validator.Validate
	Middleware *middleware.Middleware
}

func NewHandler(service Service, validator *validator.Validate, middleware *middleware.Middleware) *Handler {
	return &Handler{service: service, validator: validator, Middleware: middleware}
}

func (h *Handler) RegisterRoute(router *httprouter.Router) {
	router.POST("/api/devices", h.handleRegisterNewDevice)
	router.GET("/api/devices", h.Middleware.TimeoutMiddleware(2, h.handleGetDevices))
	router.GET("/api/devices/:device-id", h.Middleware.TimeoutMiddleware(2, h.handleGetDeviceByID))
	router.PATCH("/api/devices/:device-id", h.handleUpdate)
}

func (h *Handler) handleRegisterNewDevice(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var payload *web.DevicePayload

	if err := utils.ParseJson(r, &payload); err != nil {

		utils.JsonInternalError(w, err.Error(), nil)
		return
	}

	if err := utils.ValidateStruct(h.validator, payload); len(err) > 0 {
		utils.JsonBadRequest(w, "Required parameter is missing", err)
		return
	}
	ID, err := h.service.RegisterNewDevice(r.Context(), payload)
	if err != nil {
		switch err {
		case utils.DevicesIdAlrRegistered:
			utils.JsonConflict(w, fmt.Sprintf("Device :%s already registered,try another one", payload.DeviceId), nil)
			return
		default:
			utils.JsonInternalError(w, err.Error(), nil)
			return
		}
	}

	utils.WriteJson(w, http.StatusOK, "status ok", "Register device succesful", ID)
}

func (h *Handler) handleGetDevices(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	devices, err := h.service.GetAllDevice(r.Context())
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
	deviceID := params.ByName("device-id")
	queryParam := r.URL.Query()

	if len(queryParam) != 1 {
		utils.JsonConflict(w, "only one query parameter allowed", nil)
		return
	}

	newDeviceID := strings.TrimSpace(queryParam.Get("new-id"))
	newRoomName := strings.TrimSpace(queryParam.Get("room-name"))

	payload := &web.DevicePayload{DeviceId: newDeviceID, RoomName: newRoomName}

	if err := utils.ValidateStruct(h.validator, payload); len(err) > 0 {
		utils.JsonBadRequest(w, "Required parameter is missing", err)
		return
	}

	if newDeviceID != "" {
		newID, err := h.service.UpdateDeviceID(r.Context(), deviceID, payload)
		// log.Println("payload route:", payload)
		if err != nil {
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
		utils.WriteJson(w, 200, "status ok", "", newID)

	} else if newRoomName != "" {
		payload.DeviceId = deviceID
		newRoom, err := h.service.UpdateRoomName(r.Context(), payload)
		if err != nil {
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
		utils.WriteJson(w, 200, "status ok", "", newRoom)

	} else {
		utils.JsonBadRequest(w, "query parameter not allowed", nil)

	}

}
