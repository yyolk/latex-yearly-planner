package sections

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type DailyNotesParameters struct {
}

type DailyNotes struct {
	parameters DailyNotesParameters
	day        calendar.Day
}

func NewDailyNotes(day calendar.Day, parameters DailyNotesParameters) DailyNotes {
	return DailyNotes{
		day:        day,
		parameters: parameters,
	}
}

func (r DailyNotes) Build() ([]string, error) {
	return []string{
		r.day.Format("2006-01-02"),
	}, nil
}

func (r DailyNotes) Link(text string) string {
	return fmt.Sprintf(`\hyperlink{daily-notes-%s}{%s}`, r.day.Format("2006-01-02"), text)
}
