package sections

import (
	"errors"
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app3/components"
	"github.com/kudrykv/latex-yearly-planner/app3/types"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type MOSHeaderParameters struct {
	AfterHeaderSkip           types.Millimeters
	MonthAndQuarterSpace      types.Millimeters
	HeadingTabLineParameters  components.TabLineParameters
	QuartersTabLineParameters components.TabLineParameters
	MonthsTabLineParameters   components.TabLineParameters
}

type MOSHeaderDaily struct {
	today      calendar.Day
	tabLine    components.TabLine
	parameters MOSHeaderParameters
	year       calendar.Year
	quarter    calendar.Quarter
	month      calendar.Month
}

var ErrIncompleteDay = errors.New("incomplete day")

func NewMOSHeaderDaily(today calendar.Day, tabs components.Tabs, parameters MOSHeaderParameters) (MOSHeaderDaily, error) {
	tabLine := components.NewTabLine(tabs, parameters.HeadingTabLineParameters)

	year := today.CalendarYear()
	if year == nil {
		return MOSHeaderDaily{}, fmt.Errorf("get year: %w", ErrIncompleteDay)
	}

	quarter := today.CalendarQuarter()
	if quarter == nil {
		return MOSHeaderDaily{}, fmt.Errorf("get quarter: %w", ErrIncompleteDay)
	}

	month := today.CalendarMonth()
	if month == nil {
		return MOSHeaderDaily{}, fmt.Errorf("get month: %w", ErrIncompleteDay)
	}

	return MOSHeaderDaily{
		today:      today,
		year:       *year,
		quarter:    *quarter,
		month:      *month,
		tabLine:    tabLine,
		parameters: parameters,
	}, nil
}

func (r MOSHeaderDaily) Build() ([]string, error) {
	return []string{fmt.Sprintf(
		dailyHeaderTemplate,
		r.months().Build(),
		r.parameters.MonthAndQuarterSpace,
		r.quarters().Build(),
		r.today.Format("Monday, 2"),
		r.tabLine.Build(),
		r.parameters.AfterHeaderSkip,
	)}, nil
}

const dailyHeaderTemplate = `\marginnote{\rotatebox[origin=tr]{90}{%s\hspace{%s}%s}}%%
%s%%
\hfill{}%%
%s
\myLinePlain
\vskip%s

`

func (r MOSHeaderDaily) months() components.TabLine {
	tabs := components.Tabs{}
	months := r.year.Months()

	for i := len(months) - 1; i >= 0; i-- {
		target := months[i].Month() == r.month.Month()

		tabs = append(tabs, components.Tab{Text: months[i].Month().String()[:3], Target: target})
	}

	return components.NewTabLine(tabs, r.parameters.MonthsTabLineParameters)
}

func (r MOSHeaderDaily) quarters() components.TabLine {
	tabs := components.Tabs{}

	for i := len(r.year.Quarters) - 1; i >= 0; i-- {
		target := r.year.Quarters[i].Number() == r.quarter.Number()

		tabs = append(tabs, components.Tab{Text: r.year.Quarters[i].Name(), Target: target})
	}

	return components.NewTabLine(tabs, r.parameters.QuartersTabLineParameters)
}
