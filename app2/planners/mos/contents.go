package mos

import (
	"fmt"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type dailyContents struct {
	day calendar.Day
}

func (m dailyContents) Build() ([]string, error) {
	var hours []string

	height := `\parbox{0pt}{\vskip5mm}`

	for i := 5; i <= 23; i++ {
		strHour := fmt.Sprintf("%0d", i)
		hours = append(hours, height+strHour+`\myLineLightGray
\vskip5mm\myLineGray`)
	}

	var priorities []string

	for i := 0; i < 8; i++ {
		priorities = append(priorities, height+`$\square$\myLineGray`)
	}

	return []string{
		`\noindent\vskip1mm\begin{minipage}[t]{\myLengthThreeColumnWidth}
\myUnderline{Schedule\textcolor{white}{g}}
` + strings.Join(hours, "\n") + `
\vskip5mm\myLineLightGray
\end{minipage}%
\hspace{5mm}%
\begin{minipage}[t]{\dimexpr2\myLengthThreeColumnWidth+\myLengthThreeColumnsSeparatorWidth}
\myUnderline{Top Priorities}
` + strings.Join(priorities, "\n") + `
\vskip7mm\myUnderline{Notes}
\vspace{5mm}\hspace{.5mm}\vbox to 0mm{\myDotGrid{30}{19}}
\end{minipage}`,
	}, nil
}

type todoIndex struct{}

func (i todoIndex) Build() ([]string, error) {
	return []string{"index"}, nil
}

type todoContents struct{}

func (t todoContents) Build() ([]string, error) {
	return []string{"page with todos"}, nil
}

type notesIndex struct{}

func (r notesIndex) Build() ([]string, error) {
	return []string{"notes index"}, nil
}

type notesContents struct{}

func (r notesContents) Build() ([]string, error) {
	return []string{"notes"}, nil
}
