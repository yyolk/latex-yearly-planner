package devices

import (
	"errors"
	"fmt"
)

type Device interface {
}

var UnknownDeviceNameErr = errors.New("unknown device")

func New(deviceName string) (Device, error) {
	switch deviceName {
	case "supernote_a5x":
		return &SupernoteA5X{}, nil
	default:
		return nil, fmt.Errorf("%s: %w", deviceName, UnknownDeviceNameErr)
	}
}

type SupernoteA5X struct{}
