package components

import (
	"fmt"
	"time"
)

type MonthLink struct {
	month time.Month
}

func NewMonthLink(month time.Month) MonthLink {
	return MonthLink{month: month}
}

func (r MonthLink) String() string {
	return fmt.Sprintf(`\hyperlink{%s}{%s}`, r.month, r.month)
}
