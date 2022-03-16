package calendar

import "time"

type Day struct{ time.Time }

func (h Day) Add(days int) Day {
	return Day{Time: h.AddDate(0, 0, days)}
}

func (h Day) IsZero() bool {
	return h.Time.IsZero()
}
