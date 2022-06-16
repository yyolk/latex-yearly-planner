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
	selectedMonths  []time.Month
}

type headerOption func(*header)

func newHeader(layout common.Layout, ui mosUI, options ...headerOption) header {
	return header{layout: layout, ui: ui}.apply(options...)
}

func (r header) apply(options ...headerOption) header {
	for _, option := range options {
		option(&r)
	}

	return r
}

func headerWithYear(year calendar.Year) headerOption {
	return func(header *header) {
		header.year = year
	}
}

func headerWithLeft(left string) headerOption {
	return func(header *header) {
		header.left = left
	}
}

func headerWithRight(right []string) headerOption {
	return func(header *header) {
		if len(right) == 0 {
			return
		}

		header.right = make([]string, len(right))
		copy(header.right, right)
	}
}

func headerSelectQuarter(quarter calendar.Quarter) headerOption {
	return func(header *header) {
		header.selectedQuarter = quarter
	}
}

func headerSelectMonths(months ...time.Month) headerOption {
	return func(header *header) {
		header.selectedMonths = months
	}
}

func (r header) Build() ([]string, error) {
	return []string{
		`\marginnote{\rotatebox[origin=tr]{90}{%
\renewcommand{\arraystretch}{` + r.ui.HeaderMarginNotesArrayStretch + `}` + r.months() + `\qquad{}` + r.quarters() + `%
}}%
{\renewcommand{\arraystretch}{` + r.ui.HeaderArrayStretch + `}\begin{tabularx}{\linewidth}{@{}lY|` + strings.Join(strings.Split(strings.Repeat("r", len(r.right)), ""), "|") + `|@{}}
\Huge ` + r.left + `{\Huge\color{white}{Q}} & & ` + strings.Join(r.right, " & ") + ` \\ \hline
\end{tabularx}}

`,
	}, nil
}

func (r header) months() string {
	strs := make([]string, 0, 12)
	months := r.year.Months()

	for i := len(months) - 1; i >= 0; i-- {
		item := cell.New(months[i].Month().String()[:3])

		for _, selectedMonth := range r.selectedMonths {
			if months[i].Month() == selectedMonth {
				item = item.Ref()

				break
			}
		}

		strs = append(strs, item.Build())
	}

	return `\begin{tabularx}{` + r.ui.HeaderMarginNotesMonthsWidth + `}{*{12}{|Y}|}
	` + r.maybeHLineLeft() + strings.Join(strs, " & ") + `\\ ` + r.maybeHLineRight() + `
\end{tabularx}`
}

func (r header) quarters() string {
	quarters := make([]string, 0, 4)

	for i := 3; i >= 0; i-- {
		item := cell.New(r.year.Quarters[i].Name())

		if r.selectedQuarter.Name() == r.year.Quarters[i].Name() {
			item = item.Ref()
		}

		quarters = append(quarters, item.Build())
	}

	return `\begin{tabularx}{` + r.ui.HeaderMarginNotesQuartersWidth + `}{*{4}{|Y}|}
	` + r.maybeHLineLeft() + strings.Join(quarters, " & ") + ` \\ ` + r.maybeHLineRight() + `
\end{tabularx}`
}

func (r header) maybeHLineLeft() string {
	if r.layout.Hand == common.LeftHand {
		return `\hline{}`
	}

	return ""
}

func (r header) maybeHLineRight() string {
	if r.layout.Hand == common.RightHand {
		return `\hline{}`
	}

	return ""
}
