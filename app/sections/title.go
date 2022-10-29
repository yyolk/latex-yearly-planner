package sections

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app/calendar"
)

type TitleParameters struct {
}

type Title struct {
	parameters TitleParameters
	year       calendar.Year
}

func NewTitle(year calendar.Year, parameters TitleParameters) (Title, error) {
	return Title{
		year:       year,
		parameters: parameters,
	}, nil
}

func (r Title) Build() ([]string, error) {
	return []string{fmt.Sprintf(`{\Huge{%d}}`, r.year.Year())}, nil
}
