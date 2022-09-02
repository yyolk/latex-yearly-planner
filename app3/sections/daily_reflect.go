package sections

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type DailyReflectParameters struct {
}

type DailyReflect struct {
	day        calendar.Day
	parameters DailyReflectParameters
}

func NewDailyReflect(day calendar.Day, parameters DailyReflectParameters) DailyReflect {
	return DailyReflect{
		day:        day,
		parameters: parameters,
	}
}

func (r DailyReflect) Build() ([]string, error) {
	return []string{"reflect"}, nil
}

func (r DailyReflect) Link() string {
	return fmt.Sprintf(`\hyperlink{daily-reflect-%s}{%s}`, r.day.Format("2006-01-02"), "Reflect")
}

func (r DailyReflect) Reference() string {
	return fmt.Sprintf(`daily-reflect-%s`, r.day.Format("2006-01-02"))
}
