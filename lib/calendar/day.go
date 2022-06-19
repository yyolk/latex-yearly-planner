package calendar

import "time"

type Days []Day

type Day struct {
	moment time.Time
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
