package texcalendar_test

import (
	"io"
	"os"
	"testing"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	. "github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
	. "github.com/smartystreets/goconvey/convey"
)

func TestYear_BuildCalendar(t *testing.T) {
	t.Parallel()

	Convey("BuildCalendar", t, func() {
		Convey("when right hand is set", func() {
			Convey("should put week column to the left", func() {
				year := NewYear(2022, WithHand(common.RightHand))

				calendar := year.BuildCalendar()

				So(calendar, ShouldEqual, readFile("testdata/year_2022_right_hand.txt"))
			})
		})
	})
}

func readFile(name string) string {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	defer func() { _ = file.Close() }()

	contents, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return string(contents)
}
