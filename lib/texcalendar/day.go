package texcalendar

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Days []Day

type Day struct {
	Day calendar.Day
}

func NewDay(day calendar.Day) Day {
	return Day{Day: day}
}

func (d Day) NameAndDate() string {
	return d.Day.Format("Monday, _2")
}

func (d Day) Ref() string {
	return d.Day.Format(time.RFC3339)
}

func (d Day) Month() time.Month {
	return d.Day.Month()
}

func (d Day) Week() Week {
	return NewWeek(common.RightHand, *d.Day.Week(), false)
}
