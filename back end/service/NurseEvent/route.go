package nurseevent

import (
	"backend/middleware"
	"backend/model/web"
	"backend/utils"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	service    Service
	Validator  *validator.Validate
	Middleware *middleware.Middleware
}

func NewHandler(service Service, validator *validator.Validate, middleware *middleware.Middleware) *Handler {
	return &Handler{service: service, Validator: validator, Middleware: middleware}
}

func (h *Handler) RegisterRoute(router *httprouter.Router) {
	router.POST("/api/nurse-events", h.handleRegisterNewEvent)
	router.GET("/api/nurse-events", h.Middleware.TimeoutMiddleware(2*time.Second, h.handleGetNurseEvents))
	router.GET("/api/nurse-events/:device-id", h.Middleware.TimeoutMiddleware(2*time.Second, h.handleGetNsByID))

}

func (h *Handler) handleGetNurseEvents(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	queryParam := r.URL.Query()

	switch len(queryParam) {
	case 0:
		alerts, err := h.service.GetLastEventByID(r.Context())
		if err != nil {
			// log.Println("something went wrong", err)
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
		return
	case 1:
		logs := strings.TrimSpace(queryParam.Get("logs"))
		if logs == "true" {
			web, err := h.service.GetAllEvents(r.Context())
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
			utils.WriteJson(w, 200, "200/status ok", "", web)

		} else {
			alerts, err := h.service.GetLastEventByID(r.Context())
			if err != nil {
				// log.Println("something went wrong", err)
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
			return
		}

	default:
		log.Println("something went wrong")
		utils.JsonBadRequest(w, "query parameter not allowed", nil)

	}

}

func (h *Handler) handleGetNsByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	deviceID := params.ByName("device-id")
	nurseEvents, err := h.service.GetLatestEvent(r.Context(), deviceID)
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
		utils.JsonInternalError(w, err.Error(), nil)
		return
	}
	if err := utils.ValidateStruct(h.Validator, payload); len(err) > 0 {
		utils.WriteJson(w, http.StatusBadRequest, "bad request", "Required parameter is missing", err)
		return
	}

	err := h.service.CreateNewEvent(r.Context(), payload)
	if err != nil {
		switch err {
		case utils.NotFoundDevices:
			utils.JsonNotFound(w, "device-id not registered yet,make sure the device-id exist", nil)
			return
		default:
			log.Println("error:", err)
			utils.JsonInternalError(w, err.Error(), nil)
			return
		}
	}
	utils.WriteJson(w, 200, "status ok", payload.Event, nil)

}
