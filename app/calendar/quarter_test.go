package calendar_test

import (
	"testing"
	"time"

	calendar2 "github.com/kudrykv/latex-yearly-planner/app/calendar"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewQuarter(t *testing.T) {
	t.Parallel()

	Convey("NewQuarter", t, func() {
		qrtr := calendar2.NewQuarter(2022, calendar2.FirstQuarter, time.Monday)
		expected := calendar2.Quarter{
			Months: [3]calendar2.Month{
				januaryWithMondayFirst(),
				februaryWithMondayFirst(),
				marchWithMondayFirst(),
			},
		}

		So(qrtr, ShouldResemble, expected)
	})
}

func januaryWithMondayFirst() calendar2.Month {
	return calendar2.Month{
		Weeks: []calendar2.Week{
			calendar2.NewWeek(calendar2.FromMonth(2022, time.January, time.Monday)),
			calendar2.NewWeek(calendar2.FromTime(time.Date(2022, time.January, 3, 0, 0, 0, 0, time.Local))),
			calendar2.NewWeek(calendar2.FromTime(time.Date(2022, time.January, 10, 0, 0, 0, 0, time.Local))),
			calendar2.NewWeek(calendar2.FromTime(time.Date(2022, time.January, 17, 0, 0, 0, 0, time.Local))),
			calendar2.NewWeek(calendar2.FromTime(time.Date(2022, time.January, 24, 0, 0, 0, 0, time.Local))),
			calendar2.NewWeek(calendar2.FromWeek([7]calendar2.Day{
				calendar2.NewDay(time.Date(2022, time.January, 31, 0, 0, 0, 0, time.Local)),
				{}, {}, {}, {}, {}, {},
			})),
		},
	}
}

func marchWithMondayFirst() calendar2.Month {
	return calendar2.Month{
		Weeks: []calendar2.Week{
			calendar2.NewWeek(calendar2.FromMonth(2022, time.March, time.Monday)),
			calendar2.NewWeek(calendar2.FromTime(time.Date(2022, time.March, 7, 0, 0, 0, 0, time.Local))),
			calendar2.NewWeek(calendar2.FromTime(time.Date(2022, time.March, 14, 0, 0, 0, 0, time.Local))),
			calendar2.NewWeek(calendar2.FromTime(time.Date(2022, time.March, 21, 0, 0, 0, 0, time.Local))),
			calendar2.NewWeek(calendar2.FromWeek([7]calendar2.Day{
				calendar2.NewDay(time.Date(2022, time.March, 28, 0, 0, 0, 0, time.Local)),
				calendar2.NewDay(time.Date(2022, time.March, 29, 0, 0, 0, 0, time.Local)),
				calendar2.NewDay(time.Date(2022, time.March, 30, 0, 0, 0, 0, time.Local)),
				calendar2.NewDay(time.Date(2022, time.March, 31, 0, 0, 0, 0, time.Local)),
				{}, {}, {},
			})),
		},
	}
}
