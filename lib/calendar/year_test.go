package calendar_test

import (
	"testing"
	"time"

	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewYear(t *testing.T) {
	t.Parallel()

	Convey("NewYear", t, func() {
		year := calendar.NewYear(2022, time.Monday)
		firstQrtr := year.Quarters[calendar.FirstQuarter]
		expectedFirstQuarter := calendar.Quarter{
			Months: [3]calendar.Month{
				januaryWithMondayFirst(),
				februaryWithMondayFirst(),
				marchWithMondayFirst(),
			},
		}

		So(firstQrtr, ShouldResemble, expectedFirstQuarter)
	})
}
