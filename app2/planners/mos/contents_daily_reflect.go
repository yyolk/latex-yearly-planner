package mos

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/tex/components"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type dailyReflectContents struct {
	day texcalendar.Day
	ui  UI
}

func (r dailyReflectContents) Build() ([]string, error) {
	return []string{
		fmt.Sprintf(
			dailyReflectContentsFormat,
			r.makeGratefulSection(),
			r.makeBestThingSection(),
			r.makeDailyLogSection(),
		),
	}, nil
}

func (r dailyReflectContents) makeGratefulSection() string {
	return components.NewMesh(r.ui.ReflectGratefulRows, r.ui.ReflectGratefulCols).Build()
}

func (r dailyReflectContents) makeBestThingSection() string {
	return components.NewMesh(r.ui.ReflectBestThingRows, r.ui.ReflectBestThingCols).Build()
}

func (r dailyReflectContents) makeDailyLogSection() string {
	return components.NewMesh(r.ui.ReflectLogRows, r.ui.ReflectLogCols).Build()
}

const dailyReflectContentsFormat = `\vspace{1mm}\myUnderline{Things I'm grateful for}
%s

\vspace{18mm}\myUnderline{The best thing that has happened today}
%s

\vspace{18mm}\myUnderline{Daily log}
%s
`
