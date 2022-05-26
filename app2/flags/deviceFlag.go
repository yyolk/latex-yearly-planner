package flags

import "github.com/urfave/cli/v2"

type DeviceFlag struct {
	cli.StringFlag
}

func NewDeviceFlag() *DeviceFlag {
	return &DeviceFlag{
		StringFlag: cli.StringFlag{
			Name:     "device-name",
			Required: true,
		},
	}
}
