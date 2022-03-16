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
