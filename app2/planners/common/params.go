package common

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
)

type Params struct {
	Year     int
	Weekday  time.Weekday
	Device   devices.Device
	Sections []string
}

type ApplyParameterOption func(*Params)

func NewParams(options ...ApplyParameterOption) Params {
	params := Params{
		Weekday: time.Monday,
	}

	for _, option := range options {
		option(&params)
	}

	return params
}

func ParamWithYear(year int) ApplyParameterOption {
	return func(params *Params) {
		params.Year = year
	}
}

func ParamWithDevice(device devices.Device) ApplyParameterOption {
	return func(params *Params) {
		params.Device = device
	}
}

func ParamWithSections(sections []string) ApplyParameterOption {
	return func(params *Params) {
		params.Sections = sections
	}
}
