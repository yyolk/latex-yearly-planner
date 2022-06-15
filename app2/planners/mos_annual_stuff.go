package planners

import (
	"fmt"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app2/tex/cell"
	"github.com/kudrykv/latex-yearly-planner/app2/texsnippets"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type mosAnnualHeader struct {
	year   calendar.Year
	layout Layout

	left            string
	right           []string
	selectedQuarter calendar.Quarter
	ui              mosUI
}

type mosAnnualHeaderOption func(*mosAnnualHeader)

func newMOSAnnualHeader(layout Layout, ui mosUI, options ...mosAnnualHeaderOption) mosAnnualHeader {
	header := mosAnnualHeader{layout: layout, ui: ui}

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

func headerSelectQuarter(quarter calendar.Quarter) mosAnnualHeaderOption {
	return func(header *mosAnnualHeader) {
		header.selectedQuarter = quarter
	}
}

func (m mosAnnualHeader) Build() ([]string, error) {
	texYear := texcalendar.NewYear(m.year)

	header, err := texsnippets.Build(texsnippets.MOSHeader, map[string]string{
		"MarginNotes": `\renewcommand{\arraystretch}{2.042}` + texYear.Months() + `\qquad{}` + m.quarters(),
		"Header": `{\renewcommand{\arraystretch}{1.8185}\begin{tabularx}{\linewidth}{@{}lY|` + strings.Join(strings.Split(strings.Repeat("r", len(m.right)), ""), "|") + `|@{}}
\Huge ` + m.left + `{\Huge\color{white}{Q}} & & ` + strings.Join(m.right, " & ") + ` \\ \hline
\end{tabularx}}`,
	})

	if err != nil {
		return nil, fmt.Errorf("build header: %w", err)
	}

	return []string{header}, nil
}

func (m mosAnnualHeader) quarters() string {
	quarters := make([]string, 0, 4)

	for i := 3; i >= 0; i-- {
		item := cell.New(m.year.Quarters[i].Name())

		if m.selectedQuarter.Name() == m.year.Quarters[i].Name() {
			item = item.Ref()
		}

		quarters = append(quarters, item.Build())
	}

	return `\begin{tabularx}{5cm}{*{4}{|Y}|}
	` + strings.Join(quarters, " & ") + ` \\ \hline
\end{tabularx}`
}

type mosAnnualContents struct {
	year calendar.Year
}

func (m mosAnnualContents) Build() ([]string, error) {
	texYear := texcalendar.NewYear(m.year)

	return []string{texYear.BuildCalendar()}, nil
}
