package sections

import (
	"errors"
	"fmt"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app/components"
	"github.com/kudrykv/latex-yearly-planner/app/types"
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
	targetReference string
	linkReference   string
	today           calendar.Day
	tabLine         components.TabLine
	parameters      MOSHeaderParameters
	year            calendar.Year
	quarters        calendar.Quarters
	months          calendar.Months
	repeat          int
	titleText       string
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
		quarters:   calendar.Quarters{*quarter},
		months:     calendar.Months{*month},
		tabLine:    tabLine,
		parameters: parameters,
	}, nil
}

func NewMOSHeaderWeekly(year calendar.Year, week calendar.Week, tabs components.Tabs, parameters MOSHeaderParameters) (MOSHeaderDaily, error) {
	tabLine := components.NewTabLine(tabs, parameters.HeadingTabLineParameters)

	return MOSHeaderDaily{
		tabLine:    tabLine,
		parameters: parameters,
		year:       year,
		quarters:   week.Quarters(year.Year(), time.Monday),
		months:     week.Months(year.Year(), time.Monday),
	}, nil
}

func NewMOSHeaderMonthly(month calendar.Month, tabs components.Tabs, parameters MOSHeaderParameters) (MOSHeaderDaily, error) {
	tabLine := components.NewTabLine(tabs, parameters.HeadingTabLineParameters)

	return MOSHeaderDaily{
		tabLine:    tabLine,
		parameters: parameters,
		year:       month.Year(),
		quarters:   calendar.Quarters{month.Quarter()},
		months:     calendar.Months{month},
	}, nil
}

func NewMOSHeaderQuarterly(
	quarter calendar.Quarter, tabs components.Tabs, parameters MOSHeaderParameters,
) (MOSHeaderDaily, error) {
	tabLine := components.NewTabLine(tabs, parameters.HeadingTabLineParameters)

	return MOSHeaderDaily{
		tabLine:    tabLine,
		parameters: parameters,
		year:       quarter.Year(),
		quarters:   calendar.Quarters{quarter},
	}, nil
}

func NewMOSHeaderIncomplete(year calendar.Year, tabs components.Tabs, parameters MOSHeaderParameters) (MOSHeaderDaily, error) {
	tabLine := components.NewTabLine(tabs, parameters.HeadingTabLineParameters)

	return MOSHeaderDaily{
		tabLine:    tabLine,
		parameters: parameters,
		year:       year,
	}, nil
}

func (r MOSHeaderDaily) Build() ([]string, error) {
	repeat := r.repeat
	if repeat <= 1 {
		repeat = 1
	}

	pages := make([]string, 0, repeat)

	for i := 1; i <= repeat; i++ {
		target := r.target()
		postfix := ""
		if i > 1 {
			postfix = fmt.Sprintf(" %d", i)
			target += postfix
		}

		pages = append(pages, fmt.Sprintf(
			dailyHeaderTemplate,
			r.makeMonths().Build(),
			r.parameters.MonthAndQuarterSpace,
			r.makeQuarters().Build(),
			target,
			r.title()+postfix,
			r.tabLine.Build(),
			r.parameters.AfterHeaderSkip,
		))
	}

	return pages, nil
}

const dailyHeaderTemplate = `\marginnote{\rotatebox[origin=tr]{90}{%s\hspace{%s}%s}}%%
%s%s%%
\hfill{}%%
%s
\myLinePlain
\vskip%s

`

func (r MOSHeaderDaily) makeMonths() components.TabLine {
	tabs := components.Tabs{}
	months := r.year.Months()

	for i := len(months) - 1; i >= 0; i-- {
		target := false

		for _, month := range r.months {
			target = target || months[i].Month() == month.Month()
		}

		tabs = append(tabs, components.Tab{Text: months[i].Month().String()[:3], Target: target})
	}

	return components.NewTabLine(tabs, r.parameters.MonthsTabLineParameters)
}

func (r MOSHeaderDaily) makeQuarters() components.TabLine {
	tabs := components.Tabs{}

	for i := len(r.year.Quarters) - 1; i >= 0; i-- {
		target := false

		for _, quarter := range r.quarters {
			target = target || r.year.Quarters[i].Number() == quarter.Number()
		}

		tabs = append(tabs, components.Tab{Text: r.year.Quarters[i].Name(), Target: target})
	}

	return components.NewTabLine(tabs, r.parameters.QuartersTabLineParameters)
}

func (r MOSHeaderDaily) target() string {
	if len(r.targetReference) == 0 {
		return ""
	}

	return fmt.Sprintf(`\hypertarget{%s}{}`, r.targetReference)
}

func (r MOSHeaderDaily) title() string {
	if len(r.linkReference) == 0 {
		return r.titleText
	}

	return fmt.Sprintf(`\hyperlink{%s}{%s}`, r.linkReference, r.titleText)
}

func (r MOSHeaderDaily) Target(referencer interface{ Reference() string }) MOSHeaderDaily {
	r.targetReference = referencer.Reference()

	return r
}

func (r MOSHeaderDaily) LinkBack(referencer interface{ Reference() string }) MOSHeaderDaily {
	r.linkReference = referencer.Reference()

	return r
}

func (r MOSHeaderDaily) Repeat(repeater interface{ Repeat() int }) MOSHeaderDaily {
	r.repeat = repeater.Repeat()

	return r
}

func (r MOSHeaderDaily) Title(titler interface{ Title() string }) MOSHeaderDaily {
	r.titleText = titler.Title()

	return r
}
