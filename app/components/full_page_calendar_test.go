package components_test

import (
	"testing"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app/calendar"
	"github.com/kudrykv/latex-yearly-planner/app/components"
	. "github.com/kudrykv/latex-yearly-planner/app/test"
	"github.com/stretchr/testify/require"
)

func TestFullPageCalendar_Build(t *testing.T) {
	t.Parallel()

	t.Run("week number to the left", func(t *testing.T) {
		t.Parallel()

		month := calendar.NewMonth(2022, time.June, time.Monday)
		parameters := components.FullPageCalendarParameters{}

		fullPageCal := components.NewFullPageCalendar(month, parameters)

		actual := fullPageCal.Build()
		require.Equal(t, Fixture("full_page_cal_weeknum_left"), actual)
	})

	t.Run("week number to the right", func(t *testing.T) {
		t.Parallel()

		month := calendar.NewMonth(2022, time.June, time.Monday)
		parameters := components.FullPageCalendarParameters{WeekNumberToTheRight: true}

		fullPageCal := components.NewFullPageCalendar(month, parameters)

		actual := fullPageCal.Build()
		require.Equal(t, Fixture("full_page_cal_weeknum_right"), actual)
	})
}
