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

type stub struct{}

func (r stub) Title() string { return "title" }

func (r stub) Reference() string { return "reference" }

func TestMOSHeader_Daily_Build(t *testing.T) {
	t.Parallel()

	yearCalendar := calendar.NewYear(2020, time.Monday)
	february1st := yearCalendar.Days()[32]

	daily, err := sections.NewMOSHeaderDaily(february1st, sampleTabs(), sampleMOSParameters())
	require.NoError(t, err)

	t.Run("sample header", func(t *testing.T) {
		t.Parallel()

		build, err := daily.Build()
		require.NoError(t, err)

		require.Len(t, build, 1)
		require.Equal(t, fixture("mos_header_from_dailies"), build[0])
	})

	t.Run("the header with the title", func(t *testing.T) {
		t.Parallel()

		build, err := daily.Title(stub{}).Build()
		require.NoError(t, err)

		require.Len(t, build, 1)
		require.Contains(t, build[0], `title%`)
	})

	t.Run("the header with the target", func(t *testing.T) {
		t.Parallel()

		build, err := daily.Target(stub{}).Build()
		require.NoError(t, err)

		require.Len(t, build, 1)
		require.Contains(t, build[0], `\hypertarget{reference}{}%`)
	})

	t.Run("the header with the title and the target", func(t *testing.T) {
		t.Parallel()

		build, err := daily.Title(stub{}).Target(stub{}).Build()
		require.NoError(t, err)

		require.Len(t, build, 1)
		require.Contains(t, build[0], `\hypertarget{reference}{}title%`)
	})

	t.Run("the header with the title and the linkback", func(t *testing.T) {
		t.Parallel()

		build, err := daily.Title(stub{}).LinkBack(stub{}).Build()
		require.NoError(t, err)

		require.Len(t, build, 1)
		require.Contains(t, build[0], `\hyperlink{reference}{title}%`)
	})
}

func TestMOSHeader_Weekly_Build(t *testing.T) {
	t.Parallel()

	yearCalendar := calendar.NewYear(2022, time.Monday)

	t.Run("highlighted Q1 and January", func(t *testing.T) {
		t.Parallel()

		week2 := yearCalendar.Weeks()[2]

		weekly, err := sections.NewMOSHeaderWeekly(week2, sampleTabs(), sampleMOSParameters())
		require.NoError(t, err)

		build, err := weekly.Build()
		require.NoError(t, err)

		require.Len(t, build, 1)
		require.Equal(t, fixture("mos_header_from_weeklies"), build[0])
	})

	t.Run("highlighted Q1, both February and March", func(t *testing.T) {
		t.Parallel()

		week9 := yearCalendar.Weeks()[9]

		weekly, err := sections.NewMOSHeaderWeekly(week9, sampleTabs(), sampleMOSParameters())
		require.NoError(t, err)

		build, err := weekly.Build()
		require.NoError(t, err)

		require.Len(t, build, 1)
		require.Contains(t, build[0], `\cellcolor{black}{\textcolor{white}{Q1}`)
		require.Contains(t, build[0], `\cellcolor{black}{\textcolor{white}{Feb}`)
		require.Contains(t, build[0], `\cellcolor{black}{\textcolor{white}{Mar}`)
	})

	t.Run("highlighted Q1 and Q2, March and April", func(t *testing.T) {
		t.Parallel()

		week13 := yearCalendar.Weeks()[13]

		weekly, err := sections.NewMOSHeaderWeekly(week13, sampleTabs(), sampleMOSParameters())
		require.NoError(t, err)

		build, err := weekly.Build()
		require.NoError(t, err)

		require.Len(t, build, 1)
		require.Contains(t, build[0], `\cellcolor{black}{\textcolor{white}{Q1}`)
		require.Contains(t, build[0], `\cellcolor{black}{\textcolor{white}{Q2}`)
		require.Contains(t, build[0], `\cellcolor{black}{\textcolor{white}{Mar}`)
		require.Contains(t, build[0], `\cellcolor{black}{\textcolor{white}{Apr}`)
	})
}

func TestMOSHeader_Monthly_Build(t *testing.T) {
	t.Parallel()

	yearCalendar := calendar.NewYear(2020, time.Monday)
	february := yearCalendar.Months()[1]

	monthly, err := sections.NewMOSHeaderMonthly(february, sampleTabs(), sampleMOSParameters())
	require.NoError(t, err)

	t.Run("sample header", func(t *testing.T) {
		t.Parallel()

		build, err := monthly.Build()
		require.NoError(t, err)

		require.Len(t, build, 1)
		require.Equal(t, fixture("mos_header_from_monthlies"), build[0])
	})

	t.Run("highlights month and quarter", func(t *testing.T) {
		t.Parallel()

		build, err := monthly.Build()
		require.NoError(t, err)

		require.Len(t, build, 1)
		require.Contains(t, build[0], `\cellcolor{black}{\textcolor{white}{Feb}`)
		require.Contains(t, build[0], `\cellcolor{black}{\textcolor{white}{Q1}`)
	})
}

func TestMOSHeader_Quarterly_Build(t *testing.T) {
	t.Parallel()

	yearCalendar := calendar.NewYear(2020, time.Monday)
	q1 := yearCalendar.GetQuarters()[0]

	quarterly, err := sections.NewMOSHeaderQuarterly(q1, sampleTabs(), sampleMOSParameters())
	require.NoError(t, err)

	t.Run("sample header", func(t *testing.T) {
		t.Parallel()

		build, err := quarterly.Build()
		require.NoError(t, err)

		require.Len(t, build, 1)
		require.Equal(t, fixture("mos_header_from_quarterlies"), build[0])
	})

	t.Run("highlights quarter, no month highlights", func(t *testing.T) {
		t.Parallel()

		build, err := quarterly.Build()
		require.NoError(t, err)

		require.Len(t, build, 1)
		require.Contains(t, build[0], `\cellcolor{black}{\textcolor{white}{Q1}`)
		require.NotContains(t, build[0], `\cellcolor{black}{\textcolor{white}{Jan}`)
		require.NotContains(t, build[0], `\cellcolor{black}{\textcolor{white}{Feb}`)
		require.NotContains(t, build[0], `\cellcolor{black}{\textcolor{white}{Mar}`)
	})
}

func TestMOSHeader_Annual_Build(t *testing.T) {
	t.Parallel()

	yearCalendar := calendar.NewYear(2020, time.Monday)

	yearly, err := sections.NewMOSHeaderAnnual(yearCalendar, sampleTabs(), sampleMOSParameters())
	require.NoError(t, err)

	t.Run("sample header", func(t *testing.T) {
		t.Parallel()

		build, err := yearly.Build()
		require.NoError(t, err)

		require.Len(t, build, 1)
		require.Equal(t, fixture("mos_header_from_annual"), build[0])
	})
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
