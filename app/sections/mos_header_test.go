package sections_test

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app/components"
	"github.com/kudrykv/latex-yearly-planner/app/sections"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	"github.com/stretchr/testify/require"
)

func TestMOSHeader_Build(t *testing.T) {
	t.Parallel()

	yearCalendar := calendar.NewYear(2020, time.Monday)
	day := yearCalendar.Days()[32]

	daily, err := sections.NewMOSHeaderDaily(day, sampleTabs(), sampleMOSParameters())
	require.NoError(t, err)

	build, err := daily.Build()
	require.NoError(t, err)

	require.Len(t, build, 1)
	require.Equal(t, fixture("mos_header_from_dailies"), build[0])
}

func sampleTabs() components.Tabs {
	return components.Tabs{{Text: "tab1"}, {Text: "tab2"}, {Text: "tab3"}}
}

func sampleMOSParameters() sections.MOSHeaderParameters {
	return sections.MOSHeaderParameters{
		AfterHeaderSkip:      1,
		MonthAndQuarterSpace: 2,
		HeadingTabLineParameters: components.TabLineParameters{
			VerticalSpacing:     3,
			SpaceBetweenColumns: 4,
		},
		QuartersTabLineParameters: components.TabLineParameters{
			VerticalSpacing:     5,
			SpaceBetweenColumns: 6,
		},
		MonthsTabLineParameters: components.TabLineParameters{
			VerticalSpacing:     7,
			SpaceBetweenColumns: 8,
		},
	}
}

func fixture(filename string) string {
	pathToFile := path.Join("fixtures", fmt.Sprintf("%s.tex", filename))

	fileBytes, err := os.ReadFile(pathToFile)
	if err != nil {
		panic(fmt.Errorf("read fixture %s: %w", pathToFile, err))
	}

	return string(fileBytes)
}
