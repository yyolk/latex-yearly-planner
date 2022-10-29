package components_test

import (
	"testing"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app/calendar"
	"github.com/kudrykv/latex-yearly-planner/app/components"
	"github.com/stretchr/testify/assert"
)

func TestDayLink_String(t *testing.T) {
	t.Parallel()

	day := calendar.NewDay(time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC))

	link := components.NewDayLink(day)

	assert.Equal(t, `\hyperlink{2020-01-02}{2}`, link.String())

	link = link.Format("02")

	assert.Equal(t, `\hyperlink{2020-01-02}{02}`, link.String())
}
