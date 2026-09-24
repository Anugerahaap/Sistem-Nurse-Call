package utils

import (
	"backend/model/domain"
	"backend/model/web"
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(validate *validator.Validate, v any) map[string]any {
	errorList := make(map[string]any)
	err := validate.Struct(v)
	if err == nil {
		return errorList
	}

	var validationErrors validator.ValidationErrors

	if !errors.As(err, &validationErrors) {
		return errorList
	}

	vi := reflect.TypeOf(v).Elem()

	for _, vs := range validationErrors {
		field, _ := vi.FieldByName(vs.StructField())
		fieldName := field.Tag.Get("json")

		var errMsg string

		switch vs.Tag() {
		case "required":
			errMsg = fmt.Sprintf(
				"%v must at least got 1 character",
				fieldName,
			)

		case "oneof":
			errMsg = fmt.Sprintf(
				"only accepted one of %v got %v",
				vs.Param(),
				vs.Value(),
			)
		}

		errorList["error_"+fieldName] = errMsg
	}
	return errorList
}

func ConvertDomainDeviceIntoWeb(domain *domain.Device) *web.Device {

	return &web.Device{
		DeviceId: domain.DeviceId,
		RoomName: domain.RoomName,
		Tanggal:  domain.Tanggal,
		Waktu:    domain.Waktu,
	}
}

func ConvertDomainDevicesIntoSlices(domain []domain.Device) []web.Device {
	slices := make([]web.Device, 0, len(domain))

	for _, v := range domain {
		slices = append(slices, *ConvertDomainDeviceIntoWeb(&v))
	}
	return slices
}

func ConvertDomainNurseEventIntoWeb(domain *domain.NurseEvent) *web.NurseEvent {

	return &web.NurseEvent{
		DeviceId: domain.DeviceId,
		Event:    domain.Event,
		Tanggal:  domain.Tanggal,
		Waktu:    domain.Waktu,
	}
}

func ConvertDomainNurseEventsIntoSlices(domain []domain.NurseEvent) []web.NurseEvent {
	slices := make([]web.NurseEvent, 0, len(domain))

	for _, v := range domain {
		slices = append(slices, *ConvertDomainNurseEventIntoWeb(&v))
	}
	return slices
}
