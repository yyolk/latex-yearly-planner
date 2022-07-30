package mos

import (
	"fmt"
	"time"

	"github.com/imdario/mergo"
	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/app2/types"
)

type Parameters struct {
	year    int
	weekday time.Weekday

	ui UI
}

type UI struct {
	HeaderMarginNotesArrayStretch  string
	HeaderArrayStretch             string
	HeaderMarginNotesMonthsWidth   types.Millimeters
	HeaderMarginNotesQuartersWidth types.Millimeters

	LittleCalArrayStretch string

	LargeCalHeaderHeight       types.Millimeters
	TodosNumber                int
	FromScheduleHour           int
	ToScheduleHour             int
	HourFormat                 string
	EnableCalendarOnDailyPages bool

	DailyNotesRows int
	DailyNotesCols int

	ReflectGratefulCols  int
	ReflectGratefulRows  int
	ReflectBestThingCols int
	ReflectBestThingRows int
	ReflectLogCols       int
	ReflectLogRows       int
}

func newUI(layout common.Layout, overrides UI) (UI, error) {
	switch layout.Name {
	case "supernote_a5x":
		ui := UI{
			HeaderMarginNotesArrayStretch:  "2.042",
			HeaderMarginNotesMonthsWidth:   157,
			HeaderMarginNotesQuartersWidth: 56.05,
			HeaderArrayStretch:             "1.8185",

			LittleCalArrayStretch: "1.6",

			LargeCalHeaderHeight: 5,

			TodosNumber: 8,

			FromScheduleHour:           8,
			ToScheduleHour:             21,
			HourFormat:                 "15",
			EnableCalendarOnDailyPages: true,

			DailyNotesRows: 41,
			DailyNotesCols: 29,

			ReflectGratefulRows:  4,
			ReflectGratefulCols:  29,
			ReflectBestThingRows: 4,
			ReflectBestThingCols: 29,
			ReflectLogRows:       28,
			ReflectLogCols:       29,
		}

		if err := mergo.Merge(&ui, overrides, mergo.WithOverride); err != nil {
			return ui, fmt.Errorf("merge UI overrides: %w", err)
		}

		return ui, nil
	default:
		return UI{}, fmt.Errorf("%s: %w", layout.Name, common.UnknownDeviceErr)
	}
}
