package mos

import (
	"fmt"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/contents"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type Parameters struct {
	Year    int
	Weekday time.Weekday

	EnabledSections []string

	DailyParameters contents.DailyParameters

	calendar texcalendar.Year
}

func (r *Parameters) test() error {
	if len(r.EnabledSections) == 0 {
		return fmt.Errorf("no enabled sections")
	}

	return nil
}
