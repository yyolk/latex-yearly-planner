package mos

import (
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type dailyNotesContents struct {
	day texcalendar.Day
}

func (d dailyNotesContents) Build() ([]string, error) {
	return []string{
		`\vspace{5mm}\hspace{0.3mm}\vbox to 0mm{\myDotGrid{41}{29}}`,
	}, nil
}
