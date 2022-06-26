package texcalendar

import (
	"strconv"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app/tex"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Year struct {
	year       calendar.Year
	parameters Parameters
}

func NewYear(year int, options ...ApplyToParameters) Year {
	parameters := Parameters{}

	for _, option := range options {
		option(&parameters)
	}

	return Year{parameters: parameters, year: calendar.NewYear(year, parameters.Weekday)}
}

func (r Year) BuildCalendar() string {
	quarterRows := make([]string, 0, len(r.year.Quarters))

	for _, quarter := range r.year.GetQuarters() {
		row := NewQuarter(quarter, WithParameters(r.parameters)).Row()
		quarterRows = append(quarterRows, tex.TabularXLineWidth("@{}*{3}{X}@{}", row))
	}

	return strings.Join(quarterRows, "\n"+`\vfill`+"\n")
}

func (r Year) Months() Months {
	months := make(Months, 0, 12)

	for _, month := range r.year.Months() {
		months = append(months, NewMonth(month, WithParameters(r.parameters)))
	}

	return months
}

func (r Year) Quarters() Quarters {
	quaters := make(Quarters, 0, 4)

	for _, quarter := range r.year.GetQuarters() {
		quaters = append(quaters, NewQuarter(quarter, WithParameters(r.parameters)))
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
		weeks = append(weeks, NewWeek(week, WithParameters(r.parameters)))
	}

	return weeks
}

func (r Year) Name() string {
	return strconv.Itoa(r.year.Year())
}

func (r *Year) Apply(options ...ApplyToParameters) {
	for _, option := range options {
		option(&r.parameters)
	}
}
