package planners

import (
	"fmt"
	"strconv"

	"github.com/kudrykv/latex-yearly-planner/app2/texsnippets"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type mosWeeklyHeader struct {
	year calendar.Year
}

func (m mosWeeklyHeader) Build() ([]string, error) {
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

type mosWeeklyContents struct {
	week calendar.Week
}

func (m mosWeeklyContents) Build() ([]string, error) {
	return []string{strconv.Itoa(m.week.WeekNumber())}, nil
}
