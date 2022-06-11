package planners

import (
	"fmt"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app2/texsnippets"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type mosAnnualHeader struct {
	year   calendar.Year
	layout Layout

	left  string
	right []string
}

type mosAnnualHeaderOption func(*mosAnnualHeader)

func newMOSAnnualHeader(layout Layout, options ...mosAnnualHeaderOption) mosAnnualHeader {
	header := mosAnnualHeader{layout: layout}

	for _, option := range options {
		option(&header)
	}

	return header
}

func headerWithYear(year calendar.Year) mosAnnualHeaderOption {
	return func(header *mosAnnualHeader) {
		header.year = year
	}
}

func headerWithLeft(left string) mosAnnualHeaderOption {
	return func(header *mosAnnualHeader) {
		header.left = left
	}
}

func headerWithRight(right []string) mosAnnualHeaderOption {
	return func(header *mosAnnualHeader) {
		if len(right) == 0 {
			return
		}

		header.right = make([]string, len(right))
		copy(header.right, right)
	}
}

func (m mosAnnualHeader) Build() ([]string, error) {
	texYear := texcalendar.NewYear(m.year)

	header, err := texsnippets.Build(texsnippets.MOSHeader, map[string]string{
		"MarginNotes": `\renewcommand{\arraystretch}{2.042}` + texYear.Months() + `\qquad{}` + texYear.Quarters(),
		"Header": `{\renewcommand{\arraystretch}{1.8185}\begin{tabularx}{\linewidth}{@{}lY|` + strings.Join(strings.Split(strings.Repeat("r", len(m.right)), ""), "|") + `|@{}}
\Huge ` + m.left + `{\Huge\color{white}{Q}} & & ` + strings.Join(m.right, " & ") + ` \\ \hline
\end{tabularx}}`,
	})

	if err != nil {
		return nil, fmt.Errorf("build header: %w", err)
	}

	return []string{header}, nil
}

type mosAnnualContents struct {
	year calendar.Year
}

func (m mosAnnualContents) Build() ([]string, error) {
	texYear := texcalendar.NewYear(m.year)

	return []string{texYear.BuildCalendar()}, nil
}
