package components

import "fmt"

type TabLineParameters struct {
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
		rule,
		r.tabs,
	)
}

const tabLineTemplate = `\begin{tabular}{%s}
%s
\end{tabular}`
