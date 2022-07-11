package mos

import (
	"github.com/kudrykv/latex-yearly-planner/app2/tex/components"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type dailyNotesContents struct {
	day texcalendar.Day
	ui  ui
}

func (r dailyNotesContents) Build() ([]string, error) {
	return []string{r.mesh()}, nil
}

func (r dailyNotesContents) mesh() string {
	return components.NewMesh(r.ui.DailyNotesRows, r.ui.DailyNotesCols).Build()
}
