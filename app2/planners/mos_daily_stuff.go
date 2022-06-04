package planners

import (
	"fmt"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/texsnippets"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type mosDailyHeader struct {
	year calendar.Year
}

func (m mosDailyHeader) Build() ([]string, error) {
	texYear := texcalendar.NewYear(m.year)

	header, err := texsnippets.Build(texsnippets.MOSHeader, map[string]string{
		"MarginNotes": texYear.Months() + `\qquad{}` + texYear.Quarters(),
		"Header":      "hello world header",
	})

	if err != nil {
		return nil, fmt.Errorf("build header: %w", err)
	}

	return []string{header}, nil
}

type mosDailyContents struct {
	day calendar.Day
}

func (m mosDailyContents) Build() ([]string, error) {
	return []string{m.day.Format(time.RFC3339)}, nil
}
