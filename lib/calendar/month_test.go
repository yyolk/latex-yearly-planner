package calendar_test

import (
	"testing"
	"time"

	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewMonth(t *testing.T) {
	t.Parallel()

	Convey("NewMonth", t, func() {
		Convey("month from month", func() {
			actual := calendar.NewMonth(2022, time.February, time.Monday)
			expected := februaryWithMondayFirst()

			So(actual, ShouldResemble, expected)
		})
	})
}

func februaryWithMondayFirst() calendar.Month {
	return calendar.Month{Weeks: []calendar.Week{
		calendar.NewWeek(calendar.FromMonth(2022, time.February, time.Monday)),
		calendar.NewWeek(calendar.FromTime(time.Date(2022, time.February, 7, 0, 0, 0, 0, time.Local))),
		calendar.NewWeek(calendar.FromTime(time.Date(2022, time.February, 14, 0, 0, 0, 0, time.Local))),
		calendar.NewWeek(calendar.FromTime(time.Date(2022, time.February, 21, 0, 0, 0, 0, time.Local))),
		calendar.NewWeek(calendar.FromWeek([7]calendar.Day{
			calendar.NewDay(time.Date(2022, time.February, 28, 0, 0, 0, 0, time.Local)),
			{},
			{},
			{},
			{},
			{},
			{},
		})),
	}}
}
