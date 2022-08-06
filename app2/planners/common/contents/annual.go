package contents

import "github.com/kudrykv/latex-yearly-planner/lib/texcalendar"

type Annual struct {
	year texcalendar.Year
}

func NewAnnual(year texcalendar.Year) Annual {
	return Annual{year: year}
}

func (r Annual) Build() ([]string, error) {
	return []string{r.year.BuildCalendar()}, nil
}
