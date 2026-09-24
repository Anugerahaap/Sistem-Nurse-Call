package utils

import "errors"

// nurse-events
var (
	NotFoundNurse = errors.New("no nurse events available")
)

// devices
var (
	NotFoundDevices        = errors.New("no devices registered")
	DevicesIdAlrRegistered = errors.New("device-id already registered")
	DeviceSameParameter    = errors.New("device is the same as the current device")
)
