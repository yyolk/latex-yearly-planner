package mos

import (
	"fmt"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/types"
)

type Parameters struct {
	enabledSections []string

	year    int
	weekday time.Weekday
}

type ParametersOption func(*Parameters)

func NewParameters(options ...ParametersOption) *Parameters {
	parameters := Parameters{}

	for _, option := range options {
		option(&parameters)
	}

	return &parameters
}

func (r *Parameters) Layout(deviceName string) (types.Layout, error) {
	switch deviceName {
	case "supernote_a5x":
		return types.Layout{}, nil
	}

	return types.Layout{}, fmt.Errorf("unknown device: %s", deviceName)
}
