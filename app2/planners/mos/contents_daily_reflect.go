package mos

import "github.com/kudrykv/latex-yearly-planner/lib/calendar"

type dailyReflectContents struct {
	day calendar.Day
}

func (d dailyReflectContents) Build() ([]string, error) {
	return []string{
		`\vspace{1mm}\myUnderline{Things I'm grateful for}
\vspace{5mm}\hspace{.3mm}\vbox to 0mm{\myDotGrid{4}{29}}

\vspace{18mm}\myUnderline{The best thing that has happened today}
\vspace{5mm}\hspace{.3mm}\vbox to 0mm{\myDotGrid{4}{29}}

\vspace{18mm}\myUnderline{Daily log}
\vspace{5mm}\hspace{.3mm}\vbox to 0mm{\myDotGrid{28}{29}}`,
	}, nil
}
