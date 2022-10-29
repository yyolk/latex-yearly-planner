package sections

import (
	"errors"
	"fmt"
	"time"

	calendar2 "github.com/kudrykv/latex-yearly-planner/app/calendar"
	"github.com/kudrykv/latex-yearly-planner/app/components"
	"github.com/kudrykv/latex-yearly-planner/app/types"
)

type MOSHeaderParameters struct {
	AfterHeaderSkip           types.Millimeters
	MonthAndQuarterSpace      types.Millimeters
	HeadingTabLineParameters  components.TabLineParameters
	QuartersTabLineParameters components.TabLineParameters
	MonthsTabLineParameters   components.TabLineParameters
}

type MOSHeader struct {
	targetReference string
	linkReference   string
	today           calendar2.Day
	tabLine         components.TabLine
	parameters      MOSHeaderParameters
	year            calendar2.Year
	quarters        calendar2.Quarters
	months          calendar2.Months
	repeat          int
	titleText       string
}

var ErrIncompleteDay = errors.New("incomplete day")

func NewMOSHeaderDaily(today calendar2.Day, tabs components.Tabs, parameters MOSHeaderParameters) (MOSHeader, error) {
	tabLine := components.NewTabLine(tabs, parameters.HeadingTabLineParameters)

	year := today.CalendarYear()
	if year == nil {
		return MOSHeader{}, fmt.Errorf("get year: %w", ErrIncompleteDay)
	}

	quarter := today.CalendarQuarter()
	if quarter == nil {
		return MOSHeader{}, fmt.Errorf("get quarter: %w", ErrIncompleteDay)
	}

	month := today.CalendarMonth()
	if month == nil {
		return MOSHeader{}, fmt.Errorf("get month: %w", ErrIncompleteDay)
	}

	return MOSHeader{
		today:      today,
		year:       *year,
		quarters:   calendar2.Quarters{*quarter},
		months:     calendar2.Months{*month},
		tabLine:    tabLine,
		parameters: parameters,
	}, nil
}

func NewMOSHeaderWeekly(week calendar2.Week, tabs components.Tabs, parameters MOSHeaderParameters) (MOSHeader, error) {
	tabLine := components.NewTabLine(tabs, parameters.HeadingTabLineParameters)

	return MOSHeader{
		tabLine:    tabLine,
		parameters: parameters,
		year:       week.Year(),
		quarters:   week.Quarters(week.Year().Year(), time.Monday),
		months:     week.Months(week.Year().Year(), time.Monday),
	}, nil
}

func NewMOSHeaderMonthly(month calendar2.Month, tabs components.Tabs, parameters MOSHeaderParameters) (MOSHeader, error) {
	tabLine := components.NewTabLine(tabs, parameters.HeadingTabLineParameters)

	return MOSHeader{
		tabLine:    tabLine,
		parameters: parameters,
		year:       month.Year(),
		quarters:   calendar2.Quarters{month.Quarter()},
		months:     calendar2.Months{month},
	}, nil
}

func NewMOSHeaderQuarterly(
	quarter calendar2.Quarter, tabs components.Tabs, parameters MOSHeaderParameters,
) (MOSHeader, error) {
	tabLine := components.NewTabLine(tabs, parameters.HeadingTabLineParameters)

	return MOSHeader{
		tabLine:    tabLine,
		parameters: parameters,
		year:       quarter.Year(),
		quarters:   calendar2.Quarters{quarter},
	}, nil
}

func NewMOSHeaderAnnual(year calendar2.Year, tabs components.Tabs, parameters MOSHeaderParameters) (MOSHeader, error) {
	tabLine := components.NewTabLine(tabs, parameters.HeadingTabLineParameters)

	return MOSHeader{
		tabLine:    tabLine,
		parameters: parameters,
		year:       year,
	}, nil
}

func (r MOSHeader) Build() ([]string, error) {
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

func (r MOSHeader) makeMonths() components.TabLine {
	tabs := components.Tabs{}
	months := r.year.Months()

	for i := len(months) - 1; i >= 0; i-- {
		target := false

		for _, month := range r.months {
			target = target || months[i].Month() == month.Month()
		}

		tabs = append(tabs, components.Tab{Text: months[i].Month().String()[:3], Reference: months[i].Month().String(), Target: target})
	}

	return components.NewTabLine(tabs, r.parameters.MonthsTabLineParameters)
}

func (r MOSHeader) makeQuarters() components.TabLine {
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

func (r MOSHeader) target() string {
	if len(r.targetReference) == 0 {
		return ""
	}

	return fmt.Sprintf(`\hypertarget{%s}{}`, r.targetReference)
}

func (r MOSHeader) title() string {
	if len(r.linkReference) == 0 {
		return r.titleText
	}

	return fmt.Sprintf(`\hyperlink{%s}{%s}`, r.linkReference, r.titleText)
}

func (r MOSHeader) Target(referencer interface{ Reference() string }) MOSHeader {
	r.targetReference = referencer.Reference()

	return r
}

func (r MOSHeader) LinkBack(referencer interface{ Reference() string }) MOSHeader {
	r.linkReference = referencer.Reference()

	return r
}

func (r MOSHeader) Repeat(repeater interface{ Repeat() int }) MOSHeader {
	r.repeat = repeater.Repeat()

	return r
}

func (r MOSHeader) Title(titler interface{ Title() string }) MOSHeader {
	r.titleText = titler.Title()

	return r
}
