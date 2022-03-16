package calendar_test

import (
	"testing"
	"time"

	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewWeek(t *testing.T) {
	t.Parallel()

	Convey("NewWeek", t, func() {
		moment, err := time.Parse(time.RFC3339, "2022-02-01T00:00:00Z")
		So(err, ShouldBeNil)

		Convey("new from day", func() {
			day := calendar.Day{Time: moment}

			week := calendar.NewWeek(calendar.FromDay(day))

			for i := 0; i < 7; i++ {
				So(week.Days[i].Day(), ShouldEqual, i+1)
			}
		})

		Convey("new from time", func() {
			week := calendar.NewWeek(calendar.FromTime(moment))

			for i := 0; i < 7; i++ {
				So(week.Days[i].Day(), ShouldEqual, i+1)
			}
		})
	})
}
