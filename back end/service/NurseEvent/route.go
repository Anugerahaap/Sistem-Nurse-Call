package nurseevent

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
	Validator *validator.Validate
}

func NewHandler(service Service, validator *validator.Validate) *Handler {
	return &Handler{service: service, Validator: validator}
}

func (h *Handler) RegisterRoute(router *httprouter.Router) {
	router.POST("/api/nurse-events", h.handleRegisterNewEvent)
	router.GET("/api/nurse-events", h.handleGetNurseEvents)
	router.GET("/api/nurse-events/:device-id", h.handleGetNsByID)
	router.GET("/api/alerts", h.handleGetLastEvents)

}

func (h *Handler) handleGetNurseEvents(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	nurseEvents, err := h.service.GetNurseEvents(r.Context())
	if err != nil {
		switch err {
		case utils.NotFoundNurse:
			utils.JsonNotFound(w, "No nurse event available now ,make sure the device-id is correct or try creating one!", nil)
			return
		default:
			utils.JsonInternalError(w, err.Error(), nil)
			return
		}
	}

	utils.WriteJson(w, 200, "status ok", "", nurseEvents)

}

func (h *Handler) handleGetNsByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	deviceID := params.ByName("device-id")
	nurseEvents, err := h.service.GetByDeviceID(r.Context(), deviceID)
	if err != nil {
		switch err {
		case utils.NotFoundNurse:
			utils.JsonNotFound(w, "No nurse event available now ,make sure the device-id is correct or try creating one!", nil)
			return
		default:
			utils.JsonInternalError(w, err.Error(), nil)
			return
		}

	}

	utils.WriteJson(w, 200, "status ok", "", nurseEvents)
}

func (h *Handler) handleRegisterNewEvent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var payload *web.NurseEventCreatePayload
	if err := utils.ParseJson(r, &payload); err != nil {
		// fmt.Println("error:", err, " payload:", payload)

		utils.JsonInternalError(w, err.Error(), nil)
		return
	}
	if err := utils.ValidateStruct(h.Validator, payload); len(err) > 0 {
		utils.WriteJson(w, http.StatusBadRequest, "bad request", "Required parameter is missing", err)
		return
	}

	_, err := h.service.CreateNurseEvent(r.Context(), payload)
	if err != nil {
		switch err {
		case utils.NotFoundDevices:
			utils.JsonNotFound(w, "device-id not registered yet,make sure the device-id exist", nil)
			return
		default:
			fmt.Println("error:", err)
			utils.JsonInternalError(w, err.Error(), nil)
			return
		}
	}
	utils.WriteJson(w, 200, "status ok", "", payload.Event)

}
func (h *Handler) handleGetLastEvents(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	alerts, err := h.service.GetLastEvents(r.Context())
	if err != nil {
		switch err {
		case utils.NotFoundNurse:
			utils.JsonNotFound(w, "No nurse event available now ,make sure the device-id is correct or try creating one!", nil)
			return
		default:
			utils.JsonInternalError(w, err.Error(), nil)
			return
		}
	}

	utils.WriteJson(w, 200, "200/status ok", "", alerts)

}
