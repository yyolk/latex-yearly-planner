package sections

import (
	"errors"
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app3/components"
	"github.com/kudrykv/latex-yearly-planner/app3/types"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type MOSHeaderParameters struct {
	AfterHeaderSkip   types.Millimeters
	TabLineParameters components.TabLineParameters
}

type MOSHeaderDaily struct {
	today      calendar.Day
	tabLine    components.TabLine
	parameters MOSHeaderParameters
}

var ErrMissingQuarter = errors.New("missing quarter")

func NewMOSHeaderDaily(today calendar.Day, tabs components.Tabs, parameters MOSHeaderParameters) (MOSHeaderDaily, error) {
	tabLine := components.NewTabLine(tabs, parameters.TabLineParameters)

	quarter := today.CalendarQuarter()
	if quarter == nil {
		return MOSHeaderDaily{}, fmt.Errorf("partially initialized day: %w", ErrMissingQuarter)
	}

	return MOSHeaderDaily{
		today:      today,
		tabLine:    tabLine,
		parameters: parameters,
	}, nil
}

func (r MOSHeaderDaily) Build() ([]string, error) {
	return []string{fmt.Sprintf(
		dailyHeaderTemplate,
		r.today.Format("Monday, 2"),
		r.tabLine.Build(),
		r.parameters.AfterHeaderSkip,
	)}, nil
}

const dailyHeaderTemplate = `\marginnote{\rotatebox[origin=tr]{90}{hello world}}%%
%s%%
\hfill{}%%
%s
\myLinePlain
\vskip%s

`
