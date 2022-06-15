package mos

import (
	"strings"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/app2/tex/cell"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type header struct {
	year   calendar.Year
	layout common.Layout

	left  string
	right []string
	ui    mosUI

	selectedQuarter calendar.Quarter
	selectedMonth   time.Month
}

type mosAnnualHeaderOption func(*header)

func newMOSAnnualHeader(layout common.Layout, ui mosUI, options ...mosAnnualHeaderOption) header {
	header := header{layout: layout, ui: ui}

	for _, option := range options {
		option(&header)
	}

	return header
}

func headerWithYear(year calendar.Year) mosAnnualHeaderOption {
	return func(header *header) {
		header.year = year
	}
}

func headerWithLeft(left string) mosAnnualHeaderOption {
	return func(header *header) {
		header.left = left
	}
}

func headerWithRight(right []string) mosAnnualHeaderOption {
	return func(header *header) {
		if len(right) == 0 {
			return
		}

		header.right = make([]string, len(right))
		copy(header.right, right)
	}
}

func headerSelectQuarter(quarter calendar.Quarter) mosAnnualHeaderOption {
	return func(header *header) {
		header.selectedQuarter = quarter
	}
}

func headerSelectMonth(month time.Month) mosAnnualHeaderOption {
	return func(header *header) {
		header.selectedMonth = month
	}
}

func (m header) Build() ([]string, error) {
	return []string{
		`\marginnote{\rotatebox[origin=tr]{90}{%
\renewcommand{\arraystretch}{` + m.ui.HeaderMarginNotesArrayStretch + `}` + m.months() + `\qquad{}` + m.quarters() + `%
}}%
{\renewcommand{\arraystretch}{` + m.ui.HeaderArrayStretch + `}\begin{tabularx}{\linewidth}{@{}lY|` + strings.Join(strings.Split(strings.Repeat("r", len(m.right)), ""), "|") + `|@{}}
\Huge ` + m.left + `{\Huge\color{white}{Q}} & & ` + strings.Join(m.right, " & ") + ` \\ \hline
\end{tabularx}}

`,
	}, nil
}

func (m header) months() string {
	strs := make([]string, 0, 12)
	months := m.year.Months()

	for i := len(months) - 1; i >= 0; i-- {
		item := cell.New(months[i].Month().String()[:3])

		if months[i].Month() == m.selectedMonth {
			item = item.Ref()
		}

		strs = append(strs, item.Build())
	}

	return `\begin{tabularx}{` + m.ui.HeaderMarginNotesMonthsWidth + `}{*{12}{|Y}|}
	` + m.maybeHLineLeft() + strings.Join(strs, " & ") + `\\ ` + m.maybeHLineRight() + `
\end{tabularx}`
}

func (m header) quarters() string {
	quarters := make([]string, 0, 4)

	for i := 3; i >= 0; i-- {
		item := cell.New(m.year.Quarters[i].Name())

		if m.selectedQuarter.Name() == m.year.Quarters[i].Name() {
			item = item.Ref()
		}

		quarters = append(quarters, item.Build())
	}

	return `\begin{tabularx}{` + m.ui.HeaderMarginNotesQuartersWidth + `}{*{4}{|Y}|}
	` + m.maybeHLineLeft() + strings.Join(quarters, " & ") + ` \\ ` + m.maybeHLineRight() + `
\end{tabularx}`
}

func (m header) maybeHLineLeft() string {
	if m.layout.Hand == common.LeftHand {
		return `\hline{}`
	}

	return ""
}

func (m header) maybeHLineRight() string {
	if m.layout.Hand == common.RightHand {
		return `\hline{}`
	}

	return ""
}
