package planners

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/texsnippets"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type mosQuarterlyHeader struct {
	year calendar.Year
}

func (r mosQuarterlyHeader) Build() ([]string, error) {
	texYear := texcalendar.NewYear(r.year)

	built, err := texsnippets.Build(texsnippets.MOSHeader, map[string]string{
		"MarginNotes": texYear.Months() + `\qquad{}` + texYear.Quarters(),
		"Header":      "hello world header",
	})

	if err != nil {
		return nil, fmt.Errorf("texsnippets build: %w", err)
	}

	return []string{built}, nil
}

type mosQuarterlyContents struct {
	quarter calendar.Quarter
}

func (r mosQuarterlyContents) Build() ([]string, error) {
	return []string{texcalendar.NewQuarter(r.quarter).BuildPage()}, nil
}
