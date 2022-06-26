package mos

import (
	"fmt"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
)

type Parameters struct {
	year    int
	weekday time.Weekday

	ui ui
}

type ui struct {
	HeaderMarginNotesArrayStretch  string
	HeaderArrayStretch             string
	HeaderMarginNotesMonthsWidth   string
	HeaderMarginNotesQuartersWidth string

	LittleCalArrayStretch string
}

func newUI(device devices.Device) (ui, error) {
	switch device.(type) {
	case *devices.SupernoteA5X:
		return ui{
			HeaderMarginNotesArrayStretch:  "2.042",
			HeaderMarginNotesMonthsWidth:   "15.7cm",
			HeaderMarginNotesQuartersWidth: "5cm",
			HeaderArrayStretch:             "1.8185",

			LittleCalArrayStretch: "1.6",
		}, nil
	default:
		return ui{}, fmt.Errorf("%T: %w", device, common.UnknownDeviceTypeErr)
	}
}
