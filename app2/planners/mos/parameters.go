package mos

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
)

type Parameters struct {
	device devices.Device
	layout common.Layout

	year    int
	weekday time.Weekday

	ui mosUI
}

type mosUI struct {
	HeaderMarginNotesArrayStretch  string
	HeaderArrayStretch             string
	HeaderMarginNotesMonthsWidth   string
	HeaderMarginNotesQuartersWidth string
}
