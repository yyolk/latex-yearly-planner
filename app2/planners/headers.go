package planners

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/texsnippets"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type mosQuarterlyHeader struct {
	texYear texcalendar.Year
}

func (r mosQuarterlyHeader) Build() ([]string, error) {
	built, err := texsnippets.Build(texsnippets.MOSHeader, map[string]string{
		"MarginNotes": r.texYear.Months() + `\qquad{}` + r.texYear.Quarters(),
		"Header":      "hello world header",
	})

	if err != nil {
		return nil, fmt.Errorf("texsnippets build: %w", err)
	}

	return []string{built}, nil
}
