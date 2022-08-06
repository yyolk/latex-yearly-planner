package texcalendar

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/app2/types"
)

type Parameters struct {
	Hand       common.MainHand
	Weekday    time.Weekday
	FirstMonth time.Month

	LittleCalArrayStretch types.Spring

	LargeCalHeaderHeight types.Millimeters
}

type ApplyToParameters func(*Parameters)

func WithParameters(externalParameters Parameters) ApplyToParameters {
	return func(parameters *Parameters) {
		*parameters = externalParameters
	}
}

func WithLittleCalArrayStretch(arrayStretch types.Spring) ApplyToParameters {
	return func(parameters *Parameters) {
		parameters.LittleCalArrayStretch = arrayStretch
	}
}

func WithLargeCalHeaderHeight(largeCalHeaderHeight types.Millimeters) ApplyToParameters {
	return func(parameters *Parameters) {
		parameters.LargeCalHeaderHeight = largeCalHeaderHeight
	}
}
