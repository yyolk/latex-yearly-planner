package calendar

import "time"

type Days []Day

type Day struct {
	moment time.Time

	week    *Week
	month   *Month
	quarter *Quarter
	year    *Year
}

func NewDay(moment time.Time) Day {
	return Day{moment: moment}
}

func (r Day) Add(days int) Day {
	return Day{moment: r.moment.AddDate(0, 0, days)}
}

func (r Day) IsZero() bool {
	return r.moment.IsZero()
}

func (r Day) Weekday() time.Weekday {
	return r.moment.Weekday()
}

func (r Day) Month() time.Month {
	return r.moment.Month()
}

func (r Day) ISOWeek() (year, week int) {
	return r.moment.ISOWeek()
}

func (r Day) Year() int {
	return r.moment.Year()
}

func (r Day) Format(format string) string {
	return r.moment.Format(format)
}

func (r Day) Day() int {
	return r.moment.Day()
}

func (r Day) Week() *Week {
	return r.week
}

func (r Day) CalendarMonth() *Month {
	return r.month
}

func (r Day) CalendarQuarter() *Quarter {
	return r.quarter
}

func (r Day) CalendarYear() *Year {
	return r.year
}

func (r Day) enrich(week Week, month Month, quarter Quarter, year Year) Day {
	r.week = &week
	r.month = &month
	r.quarter = &quarter
	r.year = &year

	return r
}

func (r Day) Moment() time.Time {
	return r.moment
}
