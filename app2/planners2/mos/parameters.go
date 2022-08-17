package mos

import (
	"fmt"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/types"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type Parameters struct {
	EnabledSections []string

	calendar      texcalendar.Year
	weekday       time.Weekday
	dailySchedule types.Schedule
}

func (r *Parameters) test() error {
	if len(r.EnabledSections) == 0 {
		return fmt.Errorf("no enabled sections")
	}

	return nil
}
