package calendar_test

import (
	"testing"
	"time"

	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewQuarter(t *testing.T) {
	t.Parallel()

	Convey("NewQuarter", t, func() {
		qrtr := calendar.NewQuarter(2022, calendar.FirstQuarter, time.Monday)
		expected := calendar.Quarter{
			Months: [3]calendar.Month{
				januaryWithMondayFirst(),
				februaryWithMondayFirst(),
				marchWithMondayFirst(),
			},
		}

		So(qrtr, ShouldResemble, expected)
	})
}

func januaryWithMondayFirst() calendar.Month {
	return calendar.Month{
		Weeks: []calendar.Week{
			calendar.NewWeek(calendar.FromMonth(2022, time.January, time.Monday)),
			calendar.NewWeek(calendar.FromTime(time.Date(2022, time.January, 3, 0, 0, 0, 0, time.Local))),
			calendar.NewWeek(calendar.FromTime(time.Date(2022, time.January, 10, 0, 0, 0, 0, time.Local))),
			calendar.NewWeek(calendar.FromTime(time.Date(2022, time.January, 17, 0, 0, 0, 0, time.Local))),
			calendar.NewWeek(calendar.FromTime(time.Date(2022, time.January, 24, 0, 0, 0, 0, time.Local))),
			calendar.NewWeek(calendar.FromWeek([7]calendar.Day{
				calendar.NewDay(time.Date(2022, time.January, 31, 0, 0, 0, 0, time.Local)),
				{}, {}, {}, {}, {}, {},
			})),
		},
	}
}

func marchWithMondayFirst() calendar.Month {
	return calendar.Month{
		Weeks: []calendar.Week{
			calendar.NewWeek(calendar.FromMonth(2022, time.March, time.Monday)),
			calendar.NewWeek(calendar.FromTime(time.Date(2022, time.March, 7, 0, 0, 0, 0, time.Local))),
			calendar.NewWeek(calendar.FromTime(time.Date(2022, time.March, 14, 0, 0, 0, 0, time.Local))),
			calendar.NewWeek(calendar.FromTime(time.Date(2022, time.March, 21, 0, 0, 0, 0, time.Local))),
			calendar.NewWeek(calendar.FromWeek([7]calendar.Day{
				calendar.NewDay(time.Date(2022, time.March, 28, 0, 0, 0, 0, time.Local)),
				calendar.NewDay(time.Date(2022, time.March, 29, 0, 0, 0, 0, time.Local)),
				calendar.NewDay(time.Date(2022, time.March, 30, 0, 0, 0, 0, time.Local)),
				calendar.NewDay(time.Date(2022, time.March, 31, 0, 0, 0, 0, time.Local)),
				{}, {}, {},
			})),
		},
	}
}
