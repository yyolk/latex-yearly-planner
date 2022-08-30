package components

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app3/types"
)

type TabLineParameters struct {
	VerticalSpacing     types.Spring
	SpaceBetweenColumns types.Millimeters
}

type TabLine struct {
	tabs       Tabs
	parameters TabLineParameters
}

func NewTabLine(tabs Tabs, parameters TabLineParameters) TabLine {
	return TabLine{
		tabs:       tabs,
		parameters: parameters,
	}
}

func (r TabLine) Build() string {
	rule := fmt.Sprintf(`|*{%d}{l|}@{}`, len(r.tabs))

	return fmt.Sprintf(
		tabLineTemplate,
		r.parameters.VerticalSpacing,
		r.parameters.SpaceBetweenColumns,
		rule,
		r.tabs,
	)
}

const tabLineTemplate = `{%%
\renewcommand{\arraystretch}{%s}%%
\setlength{\tabcolsep}{%s}%%
\begin{tabular}{%s}
%s
\end{tabular}}`
