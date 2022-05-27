package texcalendar

import "github.com/kudrykv/latex-yearly-planner/lib/calendar"

type Month struct {
	month calendar.Month
}

func NewMonth(month calendar.Month) Month {
	return Month{month: month}
}

func (m Month) Tabular() string {
	return "" + m.month.Weeks[0].HeadMonth().String()
}
