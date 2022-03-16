package calendar_test

import (
	"testing"
	"time"

	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDay_Add(t *testing.T) {
	t.Parallel()

	Convey("Add", t, func() {
		moment := time.Now()

		today := calendar.Day{Time: moment}
		tomorrow := calendar.Day{Time: moment.AddDate(0, 0, 1)}

		So(today.Add(1), ShouldResemble, tomorrow)
	})
}

func TestDay_IsZero(t *testing.T) {
	t.Parallel()

	Convey("IsZero", t, func() {
		Convey("is zero", func() {
			zero := calendar.Day{}

			So(zero.IsZero(), ShouldBeTrue)
		})

		Convey("not zero", func() {
			notZero := calendar.Day{Time: time.Now()}

			So(notZero.IsZero(), ShouldBeFalse)
		})
	})
}
