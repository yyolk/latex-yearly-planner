package sections

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app3/components"
	"github.com/kudrykv/latex-yearly-planner/app3/types"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type DailyReflectParameters struct {
	BlocksSeparateSize       types.Millimeters
	GoalsNotesParameters     components.NotesParameters
	GratefulNotesParameters  components.NotesParameters
	BestThingNotesParameters components.NotesParameters
	DailyLogNotesParameters  components.NotesParameters
}

type DailyReflect struct {
	day            calendar.Day
	parameters     DailyReflectParameters
	goalsNotes     components.Notes
	gratefulNotes  components.Notes
	bestThingNotes components.Notes
	dailyLogNotes  components.Notes
}

func NewDailyReflect(day calendar.Day, parameters DailyReflectParameters) (DailyReflect, error) {
	goalsNotes, err := components.NewNotes(parameters.GoalsNotesParameters)
	if err != nil {
		return DailyReflect{}, fmt.Errorf("goals notes: %w", err)
	}

	gratefulNotes, err := components.NewNotes(parameters.GratefulNotesParameters)
	if err != nil {
		return DailyReflect{}, fmt.Errorf("grateful notes: %w", err)
	}

	bestThingNotes, err := components.NewNotes(parameters.BestThingNotesParameters)
	if err != nil {
		return DailyReflect{}, fmt.Errorf("best thing notes: %w", err)
	}

	dailyLogNotes, err := components.NewNotes(parameters.DailyLogNotesParameters)
	if err != nil {
		return DailyReflect{}, fmt.Errorf("daily log notes: %w", err)
	}

	return DailyReflect{
		parameters:     parameters,
		day:            day,
		goalsNotes:     goalsNotes,
		gratefulNotes:  gratefulNotes,
		bestThingNotes: bestThingNotes,
		dailyLogNotes:  dailyLogNotes,
	}, nil
}

func (r DailyReflect) Build() ([]string, error) {
	return []string{fmt.Sprintf(
		dailyReflectTemplate,
		r.goalsNotes.Build(),
		r.parameters.BlocksSeparateSize,
		r.gratefulNotes.Build(),
		r.parameters.BlocksSeparateSize,
		r.bestThingNotes.Build(),
		r.parameters.BlocksSeparateSize,
		r.dailyLogNotes.Build(),
	)}, nil
}

const dailyReflectTemplate = `
\myUnderline{Goals}
%s
\vskip%s
\myUnderline{Things I'm grateful for}
%s
\vskip%s
\myUnderline{The best thing that happened today}
%s
\vskip%s
\myUnderline{Daily log}
%s
`

func (r DailyReflect) Link() string {
	return fmt.Sprintf(`\hyperlink{daily-reflect-%s}{%s}`, r.day.Format("2006-01-02"), "Reflect")
}

func (r DailyReflect) Reference() string {
	return fmt.Sprintf(`daily-reflect-%s`, r.day.Format("2006-01-02"))
}
