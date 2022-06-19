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
			day := calendar.NewDay(moment)

			week := calendar.NewWeek(calendar.FromDay(day))

			for i := 0; i < 7; i++ {
				So(week.Days()[i].Day(), ShouldEqual, i+1)
			}
		})

		Convey("new from time", func() {
			week := calendar.NewWeek(calendar.FromTime(moment))

			for i := 0; i < 7; i++ {
				So(week.Days()[i].Day(), ShouldEqual, i+1)
			}
		})

		Convey("new from month", func() {
			Convey("start from sunday", func() {
				week := calendar.NewWeek(calendar.FromMonth(2022, time.February, time.Sunday))
				expected := calendar.NewWeek(calendar.FromWeek([7]calendar.Day{
					{},
					{},
					calendar.NewDay(time.Date(2022, time.February, 1, 0, 0, 0, 0, time.Local)),
					calendar.NewDay(time.Date(2022, time.February, 2, 0, 0, 0, 0, time.Local)),
					calendar.NewDay(time.Date(2022, time.February, 3, 0, 0, 0, 0, time.Local)),
					calendar.NewDay(time.Date(2022, time.February, 4, 0, 0, 0, 0, time.Local)),
					calendar.NewDay(time.Date(2022, time.February, 5, 0, 0, 0, 0, time.Local)),
				}))
				So(week, ShouldResemble, expected)
			})

			Convey("start from monday", func() {
				week := calendar.NewWeek(calendar.FromMonth(2022, time.February, time.Monday))
				expected := calendar.NewWeek(calendar.FromWeek([7]calendar.Day{
					{},
					calendar.NewDay(time.Date(2022, time.February, 1, 0, 0, 0, 0, time.Local)),
					calendar.NewDay(time.Date(2022, time.February, 2, 0, 0, 0, 0, time.Local)),
					calendar.NewDay(time.Date(2022, time.February, 3, 0, 0, 0, 0, time.Local)),
					calendar.NewDay(time.Date(2022, time.February, 4, 0, 0, 0, 0, time.Local)),
					calendar.NewDay(time.Date(2022, time.February, 5, 0, 0, 0, 0, time.Local)),
					calendar.NewDay(time.Date(2022, time.February, 6, 0, 0, 0, 0, time.Local)),
				}))
				So(week, ShouldResemble, expected)
			})

			Convey("start from friday", func() {
				week := calendar.NewWeek(calendar.FromMonth(2022, time.February, time.Friday))
				expected := calendar.NewWeek(calendar.FromWeek([7]calendar.Day{
					{},
					{},
					{},
					{},
					calendar.NewDay(time.Date(2022, time.February, 1, 0, 0, 0, 0, time.Local)),
					calendar.NewDay(time.Date(2022, time.February, 2, 0, 0, 0, 0, time.Local)),
					calendar.NewDay(time.Date(2022, time.February, 3, 0, 0, 0, 0, time.Local)),
				}))
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

func TestWeek_ZerofyMonth(t *testing.T) {
	t.Parallel()

	Convey("ZerofyMonth", t, func() {
		moment, err := time.Parse(time.RFC3339, "2022-02-24T00:00:00Z")
		So(err, ShouldBeNil)

		week := calendar.NewWeek(calendar.FromTime(moment)).ZerofyMonth(time.March)
		expected := calendar.NewWeek(calendar.FromWeek([7]calendar.Day{
			calendar.NewDay(time.Date(2022, time.February, 24, 0, 0, 0, 0, time.UTC)),
			calendar.NewDay(time.Date(2022, time.February, 25, 0, 0, 0, 0, time.UTC)),
			calendar.NewDay(time.Date(2022, time.February, 26, 0, 0, 0, 0, time.UTC)),
			calendar.NewDay(time.Date(2022, time.February, 27, 0, 0, 0, 0, time.UTC)),
			calendar.NewDay(time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC)),
			{},
			{},
		}))
		So(week, ShouldResemble, expected)
	})
}
