package mos

import (
	"time"
)

type Parameters struct {
	enabledSections []string

	year    int
	weekday time.Weekday
}
