package planners

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/texsnippets"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type mosMonthlyHeader struct {
	year calendar.Year
}

func (m mosMonthlyHeader) Build() ([]string, error) {
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

type mosMonthlyContents struct {
	month calendar.Month
}

func (m mosMonthlyContents) Build() ([]string, error) {
	return []string{m.month.Month().String()}, nil
}
