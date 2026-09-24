package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ParseJson(r *http.Request, stx any) error {
	err := json.NewDecoder(r.Body).Decode(stx)
	if err != nil {
		var typeError *json.UnmarshalTypeError

		if errors.As(err, &typeError) {
			return fmt.Errorf("invalid type for field '%s' expected %v type not %s", typeError.Field, typeError.Type, typeError.Value)
		}

		return fmt.Errorf("invalid json")
	}

	return nil
}

func WriteJson(w http.ResponseWriter, code int, status, message string, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	wr := &Response{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    v,
	}
	return json.NewEncoder(w).Encode(wr)
}

func JsonForbidden(w http.ResponseWriter, message string, v any) error {
	return WriteJson(w, http.StatusForbidden, "403/status forbidden", message, v)
}

func JsonBadRequest(w http.ResponseWriter, message string, v any) error {
	return WriteJson(w, http.StatusBadRequest, "400/status bad reqeust", message, v)
}

func JsonInternalError(w http.ResponseWriter, message string, v any) error {
	return WriteJson(w, http.StatusInternalServerError, "status internal server error", message, v)
}

func JsonUnauthorized(w http.ResponseWriter, message string, v any) error {
	return WriteJson(w, http.StatusUnauthorized, "401/status unauthorized", message, v)
}
func JsonConflict(w http.ResponseWriter, message string, v any) error {
	return WriteJson(w, http.StatusConflict, "409/status confict", message, v)
}

func JsonNotFound(w http.ResponseWriter, message string, v any) error {
	return WriteJson(w, http.StatusNotFound, "404/status not found", message, v)
}
func NotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		JsonNotFound(w, "this route is unavailable", nil)
	})
}
