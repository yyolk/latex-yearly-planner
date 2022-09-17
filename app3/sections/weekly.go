package sections

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app3/components"
	"github.com/kudrykv/latex-yearly-planner/app3/types"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type WeeklyParameters struct {
	OneColumnWidth      types.Millimeters
	TwoColumnWidth      types.Millimeters
	Separator           types.Millimeters
	OneColumnParameters components.NotesParameters
	TwoColumnParameters components.NotesParameters
}

type Weekly struct {
	parameters  WeeklyParameters
	week        calendar.Week
	oneColNotes components.Notes
	twoColNotes components.Notes
}

func NewWeekly(week calendar.Week, parameters WeeklyParameters) (Weekly, error) {
	oneColNotes, err := components.NewNotes(parameters.OneColumnParameters)
	if err != nil {
		return Weekly{}, fmt.Errorf("new notes: %w", err)
	}

	twoColNotes, err := components.NewNotes(parameters.TwoColumnParameters)
	if err != nil {
		return Weekly{}, fmt.Errorf("new notes: %w", err)
	}

	return Weekly{
		parameters:  parameters,
		week:        week,
		oneColNotes: oneColNotes,
		twoColNotes: twoColNotes,
	}, nil
}

func (w Weekly) Reference() string {
	var prefix string

	if w.week.First() {
		prefix = "first-"
	}

	return fmt.Sprintf("%sweek-%d", prefix, w.week.WeekNumber())
}

func (w Weekly) Title() string {
	return fmt.Sprintf("Week %d", w.week.WeekNumber())
}

func (w Weekly) Build() ([]string, error) {
	format := "Monday, 2"
	days := w.week.Days()

	return []string{fmt.Sprintf(
		weeklyTemplate,

		w.parameters.OneColumnWidth,
		days[0].Format(format),
		w.oneColNotes.Build(),
		w.parameters.Separator,
		w.parameters.OneColumnWidth,
		days[1].Format(format),
		w.oneColNotes.Build(),
		w.parameters.Separator,
		w.parameters.OneColumnWidth,
		days[2].Format(format),
		w.oneColNotes.Build(),

		w.parameters.OneColumnWidth,
		days[3].Format(format),
		w.oneColNotes.Build(),
		w.parameters.Separator,
		w.parameters.OneColumnWidth,
		days[4].Format(format),
		w.oneColNotes.Build(),
		w.parameters.Separator,
		w.parameters.OneColumnWidth,
		days[5].Format(format),
		w.oneColNotes.Build(),

		w.parameters.OneColumnWidth,
		days[6].Format(format),
		w.oneColNotes.Build(),
		w.parameters.Separator,
		w.parameters.TwoColumnWidth,
		w.twoColNotes.Build(),
	)}, nil
}

const weeklyTemplate = `\parbox{%s}{\myUnderline{%s}\par{}%s}%%
\hspace{%s}%%
\parbox{%s}{\myUnderline{%s}\par{}%s}%%
\hspace{%s}%%
\parbox{%s}{\myUnderline{%s}\par{}%s}%%
\vfill{}

\parbox{%s}{\myUnderline{%s}\par{}%s}%%
\hspace{%s}%%
\parbox{%s}{\myUnderline{%s}\par{}%s}%%
\hspace{%s}%%
\parbox{%s}{\myUnderline{%s}\par{}%s}%%
\vfill{}

\parbox{%s}{\myUnderline{%s}\par{}%s}%%
\hspace{%s}%%
\parbox{%s}{\myUnderline{Notes}\par{}%s}
`
