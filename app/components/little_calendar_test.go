package components_test

import (
	"testing"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app/calendar"
	"github.com/kudrykv/latex-yearly-planner/app/components"
	"github.com/kudrykv/latex-yearly-planner/app/test"
	"github.com/stretchr/testify/require"
)

func TestLittleCalendar_FromDay_Build(t *testing.T) {
	t.Parallel()

	year := calendar.NewYear(2022, time.Monday)
	february1st := year.Days()[32]

	table := map[string]struct {
		params components.LittleCalendarParameters
		want   string
	}{
		"empty": {
			params: components.LittleCalendarParameters{},
			want:   test.Fixture("littlecal_empty_params"),
		},

		"week number to the right": {
			params: components.LittleCalendarParameters{WeekNumberToTheRight: true},
			want:   test.Fixture("littlecal_weeknum_to_the_right"),
		},

		"week highlight": {
			params: components.LittleCalendarParameters{WeekHighlight: true},
			want:   test.Fixture("littlecal_week_highlight"),
		},

		"day highlight": {
			params: components.LittleCalendarParameters{DayHighlight: true},
			want:   test.Fixture("littlecal_day_highlight"),
		},
	}

	for name, tc := range table {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			littleCalendar, err := components.NewLittleCalendar(february1st, tc.params)
			require.NoError(t, err)

			actual := littleCalendar.Build()
			require.Equal(t, tc.want, actual)
		})
	}
}
