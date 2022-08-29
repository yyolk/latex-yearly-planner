package sections

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app3/types"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type MOSHeaderParameters struct {
	AfterHeaderSkip types.Millimeters
}

type MOSHeaderDaily struct {
	today      calendar.Day
	parameters MOSHeaderParameters
}

func NewMOSHeaderDaily(today calendar.Day, parameters MOSHeaderParameters) MOSHeaderDaily {
	return MOSHeaderDaily{
		today:      today,
		parameters: parameters,
	}
}

func (r MOSHeaderDaily) Build() ([]string, error) {
	return []string{fmt.Sprintf(
		dailyHeaderTemplate,
		r.today.Format("Monday, 2"),
		r.parameters.AfterHeaderSkip,
	)}, nil
}

const dailyHeaderTemplate = `%s%%
\hfill{}%%
\begin{tabular}{|*{3}{l|}@{}}
one & two & three
\end{tabular}
\myLinePlain
\vskip%s

`
