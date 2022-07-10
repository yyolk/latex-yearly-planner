package mos

import (
	"fmt"
	"time"

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

	LargeCalHeaderHeight       string
	TodosNumber                int
	FromScheduleHour           int
	ToScheduleHour             int
	HourFormat                 string
	EnableCalendarOnDailyPages bool
}

func newUI(layout common.Layout) (ui, error) {
	switch layout.Name {
	case "supernote_a5x":
		return ui{
			HeaderMarginNotesArrayStretch:  "2.042",
			HeaderMarginNotesMonthsWidth:   "15.7cm",
			HeaderMarginNotesQuartersWidth: "5.605cm",
			HeaderArrayStretch:             "1.8185",

			LittleCalArrayStretch: "1.6",

			LargeCalHeaderHeight: "5mm",

			TodosNumber: 8,

			FromScheduleHour:           8,
			ToScheduleHour:             21,
			HourFormat:                 "15",
			EnableCalendarOnDailyPages: true,
		}, nil
	default:
		return ui{}, fmt.Errorf("%s: %w", layout.Name, common.UnknownDeviceErr)
	}
}
