package domain

import "time"

// kusus untuk database
type Device struct {
	DeviceId string
	RoomName string
	Tanggal  time.Time
	Waktu    string
}

type NurseEvent struct {
	DeviceId string
	Event    string
	Tanggal  time.Time
	Waktu    string
}
