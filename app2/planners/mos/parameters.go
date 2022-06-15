package mos

import (
	"time"
)

type Parameters struct {
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
