package mos

import (
	"strings"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/app2/tex/cell"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type header struct {
	year   texcalendar.Year
	layout common.Layout

	title   string
	actions []string
	ui      mosUI

	selectedQuarter calendar.Quarter
	selectedMonths  []time.Month
	hand            common.MainHand
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

func headerWithTexYear(year texcalendar.Year) headerOption {
	return func(header *header) {
		header.year = year
	}
}

func headerWithHand(hand common.MainHand) headerOption {
	return func(header *header) {
		header.hand = hand
	}
}

func headerWithTitle(left string) headerOption {
	return func(header *header) {
		header.title = left
	}
}

func headerWithActions(actions []string) headerOption {
	return func(header *header) {
		if len(actions) == 0 {
			return
		}

		header.actions = make([]string, len(actions))
		copy(header.actions, actions)
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
	tabularFormat := `@{}lY|` + strings.Join(strings.Split(strings.Repeat("r", len(r.actions)), ""), "|") + `|@{}`

	left := `\Huge ` + r.title
	right := strings.Join(r.actions, " & ")

	if r.hand == common.LeftHand {
		tabularFormat = `@{}|` + strings.Join(strings.Split(strings.Repeat("l", len(r.actions)), ""), "|") + `Yr@{}`

		left, right = right, left
	}

	return []string{
		`\marginnote{\rotatebox[origin=tr]{90}{%
\renewcommand{\arraystretch}{` + r.ui.HeaderMarginNotesArrayStretch + `}` + r.months() + `\qquad{}` + r.quarters() + `%
}}%
{\renewcommand{\arraystretch}{` + r.ui.HeaderArrayStretch + `}\begin{tabularx}{\linewidth}{` + tabularFormat + `}
` + left + `{\Huge\color{white}{Q}} & & ` + right + ` \\ \hline
\end{tabularx}}

`,
	}, nil
}

func (r header) months() string {
	strs := make([]string, 0, 12)
	months := r.year.Months()

	for i := len(months) - 1; i >= 0; i-- {
		item := cell.New(months[i].ShortName())

		if months[i].IntersectsWith(r.selectedMonths) {
			item = item.Ref()
		}

		strs = append(strs, item.Build())
	}

	return `\begin{tabularx}{` + r.ui.HeaderMarginNotesMonthsWidth + `}{*{12}{|Y}|}
	` + r.maybeHLineLeft() + strings.Join(strs, " & ") + `\\ ` + r.maybeHLineRight() + `
\end{tabularx}`
}

func (r header) quarters() string {
	quarters := make([]string, 0, 4)

	for _, quarter := range r.year.Quarters().Reverse() {
		item := cell.New(quarter.Name())

		if quarter.Matches(texcalendar.NewQuarter(r.hand, r.selectedQuarter)) {
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
		return `\hline `
	}

	return ""
}

func (r header) maybeHLineRight() string {
	if r.layout.Hand == common.RightHand {
		return `\hline`
	}

	return ""
}
