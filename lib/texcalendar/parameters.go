package texcalendar

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
)

type Parameters struct {
	Hand       common.MainHand
	Weekday    time.Weekday
	FirstMonth time.Month
	ForLarge   bool
}

type ApplyToParameters func(*Parameters)

func WithParameters(externalParameters Parameters) ApplyToParameters {
	return func(parameters *Parameters) {
		*parameters = externalParameters
	}
}
