package devices

import (
	"errors"
	"fmt"
)

type Device interface {
	Paper() Paper
}

type Paper struct {
	Width  string
	Height string
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

func (s *SupernoteA5X) Paper() Paper {
	return Paper{
		Width:  "15.6cm",
		Height: "23cm",
	}
}
