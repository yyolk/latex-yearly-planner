package calendar_test

import (
	"testing"
	"time"

	calendar2 "github.com/kudrykv/latex-yearly-planner/app/calendar"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewYear(t *testing.T) {
	t.Parallel()

	Convey("NewYear", t, func() {
		year := calendar2.NewYear(2022, time.Monday)
		firstQrtr := year.Quarters[calendar2.FirstQuarter]
		expectedFirstQuarter := calendar2.Quarter{
			Months: [3]calendar2.Month{
				januaryWithMondayFirst(),
				februaryWithMondayFirst(),
				marchWithMondayFirst(),
			},
		}

		So(firstQrtr, ShouldResemble, expectedFirstQuarter)
	})
}
