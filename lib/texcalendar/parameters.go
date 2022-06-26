package texcalendar

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
)

type Parameters struct {
	Hand         common.MainHand
	Weekday      time.Weekday
	FirstMonth   time.Month
	ArrayStretch string
}

type ApplyToParameters func(*Parameters)

func WithParameters(externalParameters Parameters) ApplyToParameters {
	return func(parameters *Parameters) {
		*parameters = externalParameters
	}
}

func WithArrayStretch(arrayStretch string) ApplyToParameters {
	return func(parameters *Parameters) {
		parameters.ArrayStretch = arrayStretch
	}
}
