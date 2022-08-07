package mos

import (
	"fmt"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/types"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type Parameters struct {
	enabledSections []string

	calendar      texcalendar.Year
	weekday       time.Weekday
	dailySchedule types.Schedule
}

type ParametersOption func(*Parameters)

func NewParameters(options ...ParametersOption) *Parameters {
	parameters := Parameters{}

	for _, option := range options {
		option(&parameters)
	}

	return &parameters
}

func WithYear(year int) ParametersOption {
	return func(parameters *Parameters) {
		parameters.calendar = texcalendar.NewYear(year)
	}
}

func WithSections(sections []string) ParametersOption {
	return func(parameters *Parameters) {
		parameters.enabledSections = sections
	}
}

func (r *Parameters) Layout(deviceName string) (types.Layout, error) {
	switch deviceName {
	case "supernote_a5x":
		return types.Layout{
			Paper:  types.Paper{Width: 156, Height: 230},
			Margin: types.Margin{Top: 10, Right: 5, Bottom: 10, Left: 5},
			Sizes: types.Sizes{
				TwoColumnsSeparatorSize:   5,
				ThreeColumnsSeparatorSize: 5,
			},
			Debug: types.Debug{
				ShowLinks:  true,
				ShowFrames: true,
			},
			Misc: *r,
		}, nil
	}

	return types.Layout{}, fmt.Errorf("unknown device: %s", deviceName)
}

func (r *Parameters) test() error {
	if len(r.enabledSections) == 0 {
		return fmt.Errorf("no enabled sections")
	}

	return nil
}
