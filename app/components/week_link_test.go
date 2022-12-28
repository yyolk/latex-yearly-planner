package components_test

import (
	"testing"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app/calendar"
	"github.com/kudrykv/latex-yearly-planner/app/components"
	"github.com/stretchr/testify/assert"
)

func TestWeekLink_String(t *testing.T) {
	t.Parallel()

	moment := time.Date(2022, time.December, 26, 0, 0, 0, 0, time.UTC)
	week := calendar.NewWeek(calendar.FromTime(moment))

	t.Run("'first' flag is set", func(t *testing.T) {
		t.Parallel()

		link := components.NewWeekLink(week.SetFirst())

		assert.Equal(t, `\hyperlink{first-week-52}{Week 52}`, link.String())
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()

		link := components.NewWeekLink(week)

		assert.Equal(t, `\hyperlink{week-52}{Week 52}`, link.String())
	})

	t.Run("string with custom format", func(t *testing.T) {
		t.Parallel()

		link := components.NewWeekLink(week).Format("W%02d")

		assert.Equal(t, `\hyperlink{week-52}{W52}`, link.String())
	})
}
