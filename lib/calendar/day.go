package calendar

import "time"

type Days []Day

func (d Days) BuildLittle() []string {
	out := make([]string, 0, len(d))

	for _, day := range d {
		if day.IsZero() {
			out = append(out, "")

			continue
		}

		out = append(out, day.Format("_2"))
	}

	return out
}

func (d Days) BuildLarge() []string {
	out := make([]string, 0, len(d))

	for _, day := range d {
		if day.IsZero() {
			out = append(out, "")

			continue
		}

		name := `{\renewcommand{\arraystretch}{1.2}\begin{tabular}{@{}p{5mm}@{}|}\hfil{}` + day.Format("_2") + `\\ \hline\end{tabular}}`
		out = append(out, name)
	}

	return out
}

type Day struct {
	moment time.Time

	week    *Week
	month   *Month
	quarter *Quarter
}

func NewDay(moment time.Time) Day {
	return Day{moment: moment}
}

func (h Day) Add(days int) Day {
	return Day{moment: h.moment.AddDate(0, 0, days)}
}

func (h Day) IsZero() bool {
	return h.moment.IsZero()
}

func (h Day) Weekday() time.Weekday {
	return h.moment.Weekday()
}

func (h Day) Month() time.Month {
	return h.moment.Month()
}

func (h Day) ISOWeek() (year, week int) {
	return h.moment.ISOWeek()
}

func (h Day) Year() int {
	return h.moment.Year()
}

func (h Day) Format(format string) string {
	return h.moment.Format(format)
}

func (h Day) Day() int {
	return h.moment.Day()
}

func (h Day) Week() *Week {
	return h.week
}

func (h Day) enrich(week Week, month Month, quarter Quarter) Day {
	h.week = &week
	h.month = &month
	h.quarter = &quarter

	return h
}
