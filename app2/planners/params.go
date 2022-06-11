package planners

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
)

type Params struct {
	year     int
	weekday  time.Weekday
	device   devices.Device
	sections []string
}

type ApplyParameterOption func(*Params)

func NewParams(options ...ApplyParameterOption) Params {
	params := Params{
		weekday: time.Monday,
	}

	for _, option := range options {
		option(&params)
	}

	return params
}

func ParamWithYear(year int) ApplyParameterOption {
	return func(params *Params) {
		params.year = year
	}
}

func ParamWithDevice(device devices.Device) ApplyParameterOption {
	return func(params *Params) {
		params.device = device
	}
}

func ParamWithSections(sections []string) ApplyParameterOption {
	return func(params *Params) {
		params.sections = sections
	}
}
