package sections

import "github.com/kudrykv/latex-yearly-planner/lib/calendar"

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
