package components_test

import (
	"testing"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app/components"
	. "github.com/kudrykv/latex-yearly-planner/app/test"
	"github.com/stretchr/testify/assert"
)

func TestSchedule_Build(t *testing.T) {
	t.Parallel()

	moment := time.Date(2022, time.June, 20, 0, 0, 0, 0, time.UTC)
	parameters := components.ScheduleParameters{
		FromHour:   10,
		ToHour:     18,
		Format:     "15",
		LineHeight: 5,
	}

	t.Run("build", func(t *testing.T) {
		t.Parallel()

		schedule, err := components.NewSchedule(moment, parameters)

		assert.NoError(t, err)
		assert.Equal(t, Fixture("schedule"), schedule.Build())
	})
}
