package texcalendar

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
)

type Parameters struct {
	Hand       common.MainHand
	Weekday    time.Weekday
	FirstMonth time.Month

	LittleCalArrayStretch string

	LargeCalHeaderHeight common.Millimeters
}

type ApplyToParameters func(*Parameters)

func WithParameters(externalParameters Parameters) ApplyToParameters {
	return func(parameters *Parameters) {
		*parameters = externalParameters
	}
}

func WithLittleCalArrayStretch(arrayStretch string) ApplyToParameters {
	return func(parameters *Parameters) {
		parameters.LittleCalArrayStretch = arrayStretch
	}
}

func WithLargeCalHeaderHeight(largeCalHeaderHeight common.Millimeters) ApplyToParameters {
	return func(parameters *Parameters) {
		parameters.LargeCalHeaderHeight = largeCalHeaderHeight
	}
}
