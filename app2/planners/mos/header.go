package mos

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/pages"
	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/app2/tex/cell"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type header struct {
	year   texcalendar.Year
	layout common.Layout

	title string
	ui    UI

	selectedQuarter texcalendar.Quarter
	selectedMonths  []time.Month
	hand            common.MainHand
	action          cell.Cells
	repeat          int
}

func newHeader(layout common.Layout, ui UI, options ...headerOption) header {
	return header{layout: layout, ui: ui}.apply(options...)
}

func (r header) apply(options ...headerOption) header {
	for _, option := range options {
		option(&r)
	}

	return r
}

func (r header) Build() ([]string, error) {
	headerText, err := r.buildHeader()
	if err != nil {
		return nil, fmt.Errorf("build header: %w", err)
	}

	return repeat(headerText, r.repeat), nil
}

func (r header) buildHeader() (string, error) {
	buffer := bytes.NewBuffer(nil)

	left := `\Huge{}` + r.title + `{\color{white}{Q}}`
	right := strings.Join(r.action.Slice(), " & ")

	if r.hand == common.LeftHand {
		left = right
		right = `\Huge{\color{white}{Q}}` + r.title
	}

	err := template.
		Must(template.New("header").Parse(headerTemplate)).
		Execute(buffer, map[string]interface{}{
			"Left":                          left,
			"Right":                         right,
			"Months":                        r.months(),
			"Quarters":                      r.quarters(),
			"TabularFormat":                 r.makeTabularFormat(),
			"HeaderMarginNotesArrayStretch": r.ui.HeaderMarginNotesArrayStretch,
			"HeaderArrayStretch":            r.ui.HeaderArrayStretch,
		})

	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

const headerTemplate = `\marginnote{\rotatebox[origin=tr]{90}{%
\renewcommand{\arraystretch}{ {{- .HeaderMarginNotesArrayStretch -}} }{{ .Months }}\qquad{}{{ .Quarters }}%
}}%
{\renewcommand{\arraystretch}{ {{- .HeaderArrayStretch -}} }\begin{tabularx}{\linewidth}{ {{- .TabularFormat -}} }
{{ .Left }} & & {{ .Right }} \\ \hline
\end{tabularx}}

`

func repeat(text string, repeat int) []string {
	out := []string{text}

	for i := 1; i < repeat; i++ {
		out = append(out, text)
	}

	return out
}

func (r header) makeTabularFormat() string {
	format := `@{}lY*{%d}{r|}@{}`

	if r.hand == common.LeftHand {
		format = `@{}lY*{%d}{r|}@{}`
	}

	return fmt.Sprintf(format, len(r.action))
}

func (r header) months() string {
	strs := make([]string, 0, 12)
	months := r.year.Months()

	for i := len(months) - 1; i >= 0; i-- {
		item := cell.New(months[i].ShortName()).RefAs(months[i].Name())

		if months[i].IntersectsWith(r.selectedMonths) {
			item = item.NoTarget()
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

		if quarter.Matches(r.selectedQuarter) {
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

func (r header) repeatTimes(repeat int) pages.Block {
	r.repeat = repeat

	return r
}
