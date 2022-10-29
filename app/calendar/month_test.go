package calendar_test

import (
	"testing"
	"time"

	calendar2 "github.com/kudrykv/latex-yearly-planner/app/calendar"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewMonth(t *testing.T) {
	t.Parallel()

	Convey("NewMonth", t, func() {
		Convey("month from month", func() {
			actual := calendar2.NewMonth(2022, time.February, time.Monday)
			expected := februaryWithMondayFirst()

			So(actual, ShouldResemble, expected)
		})
	})
}

func februaryWithMondayFirst() calendar2.Month {
	return calendar2.Month{Weeks: []calendar2.Week{
		calendar2.NewWeek(calendar2.FromMonth(2022, time.February, time.Monday)),
		calendar2.NewWeek(calendar2.FromTime(time.Date(2022, time.February, 7, 0, 0, 0, 0, time.Local))),
		calendar2.NewWeek(calendar2.FromTime(time.Date(2022, time.February, 14, 0, 0, 0, 0, time.Local))),
		calendar2.NewWeek(calendar2.FromTime(time.Date(2022, time.February, 21, 0, 0, 0, 0, time.Local))),
		calendar2.NewWeek(calendar2.FromWeek([7]calendar2.Day{
			calendar2.NewDay(time.Date(2022, time.February, 28, 0, 0, 0, 0, time.Local)),
			{},
			{},
			{},
			{},
			{},
			{},
		})),
	}}
}
