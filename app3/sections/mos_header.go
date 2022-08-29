package sections

import (
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

func NewMOSHeaderDaily(today calendar.Day, tabs components.Tabs, parameters MOSHeaderParameters) MOSHeaderDaily {
	tabLine := components.NewTabLine(tabs, parameters.TabLineParameters)

	return MOSHeaderDaily{
		today:      today,
		tabLine:    tabLine,
		parameters: parameters,
	}
}

func (r MOSHeaderDaily) Build() ([]string, error) {
	return []string{fmt.Sprintf(
		dailyHeaderTemplate,
		r.today.Format("Monday, 2"),
		r.tabLine.Build(),
		r.parameters.AfterHeaderSkip,
	)}, nil
}

const dailyHeaderTemplate = `%s%%
\hfill{}%%
%s
\myLinePlain
\vskip%s

`
