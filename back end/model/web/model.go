package web

import (
	"context"
	"time"
)

// kusus untuk api
type Device struct {
	DeviceId string    `json:"device-id"`
	RoomName string    `json:"room-name"`
	Tanggal  time.Time `json:"tanggal"`
	Waktu    string    `json:"waktu"`
}

type NurseEvent struct {
	DeviceId string    `json:"device-id"`
	Event    string    `json:"event"`
	Tanggal  time.Time `json:"tanggal"`
	Waktu    string    `json:"waktu"`
}

type DevicePayload struct {
	DeviceId string `json:"device-id" validate:"required"`
	RoomName string `json:"room-name"`
}

type NurseEventCreatePayload struct {
	DeviceId string `json:"device-id" validate:"required"`
	Event    string `json:"event" validate:"required,oneof=INFUS DARURAT RESET"`
}

type DeviceServices interface {
	GetByDeviceId(ctx context.Context, deviceId string) (*Device, error)
	Create(ctx context.Context, dvs *DevicePayload) error
	Update(ctx context.Context, dvs *DevicePayload) error
	Delete(ctx context.Context, deviceId string) error
}

type NurseEventServices interface {
	Create(ctx context.Context, ns *NurseEventCreatePayload) (string, error)
	GetByDeviceId(ctx context.Context, deviceId string) ([]NurseEvent, error)
	GetNurseEvents(ctx context.Context) ([]NurseEvent, error)
}
