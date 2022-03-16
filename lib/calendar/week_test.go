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

		Convey("new from month", func() {
			Convey("start from sunday", func() {
				week := calendar.NewWeek(calendar.FromMonth(2022, time.February, time.Sunday))
				expected := calendar.Week{Days: [7]calendar.Day{
					{Time: time.Time{}},
					{Time: time.Time{}},
					{Time: time.Date(2022, time.February, 1, 0, 0, 0, 0, time.Local)},
					{Time: time.Date(2022, time.February, 2, 0, 0, 0, 0, time.Local)},
					{Time: time.Date(2022, time.February, 3, 0, 0, 0, 0, time.Local)},
					{Time: time.Date(2022, time.February, 4, 0, 0, 0, 0, time.Local)},
					{Time: time.Date(2022, time.February, 5, 0, 0, 0, 0, time.Local)},
				}}

				So(week, ShouldResemble, expected)
			})

			Convey("start from monday", func() {
				week := calendar.NewWeek(calendar.FromMonth(2022, time.February, time.Monday))
				expected := calendar.Week{Days: [7]calendar.Day{
					{Time: time.Time{}},
					{Time: time.Date(2022, time.February, 1, 0, 0, 0, 0, time.Local)},
					{Time: time.Date(2022, time.February, 2, 0, 0, 0, 0, time.Local)},
					{Time: time.Date(2022, time.February, 3, 0, 0, 0, 0, time.Local)},
					{Time: time.Date(2022, time.February, 4, 0, 0, 0, 0, time.Local)},
					{Time: time.Date(2022, time.February, 5, 0, 0, 0, 0, time.Local)},
					{Time: time.Date(2022, time.February, 6, 0, 0, 0, 0, time.Local)},
				}}

				So(week, ShouldResemble, expected)
			})

			Convey("start from friday", func() {
				week := calendar.NewWeek(calendar.FromMonth(2022, time.February, time.Friday))
				expected := calendar.Week{Days: [7]calendar.Day{
					{Time: time.Time{}},
					{Time: time.Time{}},
					{Time: time.Time{}},
					{Time: time.Time{}},
					{Time: time.Date(2022, time.February, 1, 0, 0, 0, 0, time.Local)},
					{Time: time.Date(2022, time.February, 2, 0, 0, 0, 0, time.Local)},
					{Time: time.Date(2022, time.February, 3, 0, 0, 0, 0, time.Local)},
				}}

				So(week, ShouldResemble, expected)
			})
		})
	})
}

func TestWeek_Next(t *testing.T) {
	t.Parallel()

	Convey("Next", t, func() {
		moment, err := time.Parse(time.RFC3339, "2022-02-01T00:00:00Z")
		So(err, ShouldBeNil)

		momentInWeek, err := time.Parse(time.RFC3339, "2022-02-08T00:00:00Z")
		So(err, ShouldBeNil)

		week := calendar.NewWeek(calendar.FromTime(moment))
		expected := calendar.NewWeek(calendar.FromTime(momentInWeek))

		So(week.Next(), ShouldResemble, expected)
	})
}

func TestWeek_HeadMonth(t *testing.T) {
	t.Parallel()

	Convey("HeadMonth", t, func() {
		moment, err := time.Parse(time.RFC3339, "2022-02-24T00:00:00Z")
		So(err, ShouldBeNil)

		week := calendar.NewWeek(calendar.FromTime(moment))

		So(week.HeadMonth(), ShouldEqual, time.February)
	})
}

func TestWeek_TailMonth(t *testing.T) {
	t.Parallel()

	Convey("TailMonth", t, func() {
		moment, err := time.Parse(time.RFC3339, "2022-02-24T00:00:00Z")
		So(err, ShouldBeNil)

		week := calendar.NewWeek(calendar.FromTime(moment))

		So(week.TailMonth(), ShouldEqual, time.March)
	})
}
