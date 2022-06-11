package planners

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
)

type MonthsOnSidesDetails struct {
	device devices.Device
	layout Layout

	year    int
	weekday time.Weekday
}
