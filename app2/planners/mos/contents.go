package mos

import (
	"fmt"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/app2/tex/ref"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type dailyContents struct {
	day  texcalendar.Day
	hand common.MainHand
}

func (m dailyContents) Build() ([]string, error) {
	leftColumn := m.scheduleColumn()
	rightColumn := m.prioritiesAndNotesColumn()

	if m.hand == common.LeftHand {
		leftColumn, rightColumn = rightColumn, leftColumn
	}

	return []string{
		`\noindent\vskip1mm` + leftColumn + `\hspace{5mm}` + rightColumn,
	}, nil
}

func (m dailyContents) prioritiesAndNotesColumn() string {
	var priorities []string

	for i := 0; i < 8; i++ {
		priorities = append(priorities, m.height()+`$\square$\myLineGray`)
	}

	moreNotes := ref.NewLinkWithRef("More", m.day.Ref()+"-notes").Build()
	dailyReflect := ref.NewLinkWithRef("Reflect", m.day.Ref()+"-reflect").Build()

	return `\begin{minipage}[t]{\dimexpr2\myLengthThreeColumnWidth+\myLengthThreeColumnsSeparatorWidth}
\myUnderline{Top Priorities}
` + strings.Join(priorities, "\n") + `
\vskip7mm\myUnderline{Notes | ` + moreNotes + ` \hfill{}` + dailyReflect + `}
\vspace{5mm}\hspace{.5mm}\vbox to 0mm{\myDotGrid{30}{19}}
\end{minipage}`
}

func (m dailyContents) scheduleColumn() string {
	var hours []string

	for i := 5; i <= 23; i++ {
		strHour := fmt.Sprintf("%0d", i)
		hours = append(hours, m.height()+strHour+`\myLineLightGray
\vskip5mm\myLineGray`)
	}

	return `\begin{minipage}[t]{\myLengthThreeColumnWidth}
\myUnderline{Schedule\textcolor{white}{g}}
` + strings.Join(hours, "\n") + `
\vskip5mm\myLineLightGray
\end{minipage}`
}

func (m dailyContents) height() string {
	return `\parbox{0pt}{\vskip5mm}`
}

type notesContents struct{}

func (r notesContents) Build() ([]string, error) {
	return []string{`\vskip5mm\hspace{0.5mm}\vbox to 0mm{\myDotGrid{41}{29}}`}, nil
}

type titleContents struct {
	title string
}

func (r titleContents) Build() ([]string, error) {
	return []string{`\hspace{0pt}\vfil
\hfill\resizebox{.7\linewidth}{!}{` + r.title + `}%
\pagebreak`}, nil
}
