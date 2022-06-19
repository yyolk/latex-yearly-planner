package texcalendar

import (
	"strconv"
	"strings"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Year struct {
	year calendar.Year
	hand common.MainHand
}

func NewYear(hand common.MainHand, year calendar.Year) Year {
	return Year{hand: hand, year: year}
}

func NewYearFromInt(hand common.MainHand, year int, weekday time.Weekday) Year {
	calYear := calendar.NewYear(year, weekday)

	return NewYear(hand, calYear)
}

func (r Year) BuildCalendar() string {
	quarterRows := make([]string, 0, len(r.year.Quarters))

	tabular := func(str string) string {
		return `\begin{tabularx}{\linewidth}{@{}*{3}{X}@{}}` + "\n" +
			str +
			"\n" + `\end{tabularx}`
	}

	for _, quarter := range r.year.Quarters {
		quarterRows = append(quarterRows, tabular(NewQuarter(r.hand, quarter).Row()))
	}

	return strings.Join(quarterRows, "\n"+`\vfill`+"\n")
}

func (r Year) Months() Months {
	months := make(Months, 0, 12)

	for _, month := range r.year.Months() {
		months = append(months, NewMonth(month, r.hand))
	}

	return months
}

func (r Year) Quarters() Quarters {
	quaters := make(Quarters, 0, 4)

	for _, quarter := range r.year.GetQuarters() {
		quaters = append(quaters, NewQuarter(r.hand, quarter))
	}

	return quaters
}

func (r Year) Days() Days {
	days := make(Days, 0, 366)

	for _, day := range r.year.Days() {
		days = append(days, NewDay(day))
	}

	return days
}

func (r Year) InWeeks() []Week {
	weeks := make([]Week, 0, 53)

	for _, week := range r.year.InWeeks() {
		weeks = append(weeks, NewWeek(r.hand, week, false))
	}

	return weeks
}

func (r Year) Name() string {
	return strconv.Itoa(r.year.Year())
}
