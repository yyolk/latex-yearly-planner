package sections

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app3/components"
	"github.com/kudrykv/latex-yearly-planner/app3/types"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type WeeklyParameters struct {
	LeftColumnWidth        types.Millimeters
	CenterColumnWidth      types.Millimeters
	RightColumnWidth       types.Millimeters
	TwoColumnWidth         types.Millimeters
	Separator              types.Millimeters
	LeftColumnParameters   components.NotesParameters
	CenterColumnParameters components.NotesParameters
	RightColumnParameters  components.NotesParameters
	TwoColumnParameters    components.NotesParameters
}

type Weekly struct {
	parameters     WeeklyParameters
	week           calendar.Week
	leftColNotes   components.Notes
	centerColNotes components.Notes
	rightColNotes  components.Notes
	twoColNotes    components.Notes
}

func NewWeekly(week calendar.Week, parameters WeeklyParameters) (Weekly, error) {
	leftColNotes, err := components.NewNotes(parameters.LeftColumnParameters)
	if err != nil {
		return Weekly{}, fmt.Errorf("new notes: %w", err)
	}

	centerColNotes, err := components.NewNotes(parameters.CenterColumnParameters)
	if err != nil {
		return Weekly{}, fmt.Errorf("new notes: %w", err)
	}

	rightColNotes, err := components.NewNotes(parameters.RightColumnParameters)
	if err != nil {
		return Weekly{}, fmt.Errorf("new notes: %w", err)
	}

	twoColNotes, err := components.NewNotes(parameters.TwoColumnParameters)
	if err != nil {
		return Weekly{}, fmt.Errorf("new notes: %w", err)
	}

	return Weekly{
		parameters:     parameters,
		week:           week,
		leftColNotes:   leftColNotes,
		centerColNotes: centerColNotes,
		rightColNotes:  rightColNotes,
		twoColNotes:    twoColNotes,
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
	link := func(i int) string {
		return fmt.Sprintf(`\hyperlink{%s}{%s}`, days[i].Format("2006-01-02"), days[i].Format(format))
	}

	return []string{fmt.Sprintf(
		weeklyTemplate,

		w.parameters.LeftColumnWidth,
		link(0),
		w.leftColNotes.Build(),
		w.parameters.Separator,
		w.parameters.CenterColumnWidth,
		link(1),
		w.centerColNotes.Build(),
		w.parameters.Separator,
		w.parameters.RightColumnWidth,
		link(2),
		w.rightColNotes.Build(),

		w.parameters.LeftColumnWidth,
		link(3),
		w.leftColNotes.Build(),
		w.parameters.Separator,
		w.parameters.CenterColumnWidth,
		link(4),
		w.centerColNotes.Build(),
		w.parameters.Separator,
		w.parameters.RightColumnWidth,
		link(5),
		w.rightColNotes.Build(),

		w.parameters.LeftColumnWidth,
		link(6),
		w.leftColNotes.Build(),
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
