package flags

import (
	"errors"
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
	"github.com/urfave/cli/v2"
)

type DeviceFlag struct {
	cli.StringFlag
}

var UnknownDeviceNameErr = errors.New("unknown device")

func (r DeviceFlag) Device() (devices.Device, error) {
	switch r.GetValue() {
	case "supernote_a5x":
		return &devices.SupernoteA5X{}, nil
	default:
		return nil, fmt.Errorf("%s: %w", r.GetValue(), UnknownDeviceNameErr)
	}
}

func NewDeviceFlag() *DeviceFlag {
	return &DeviceFlag{
		StringFlag: cli.StringFlag{
			Name:     "device-name",
			Required: true,
		},
	}
}
