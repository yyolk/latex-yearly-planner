package calendar

import "time"

type Year struct {
	Quarters [4]Quarter
}

func NewYear(year int, wd time.Weekday) Year {
	return Year{
		Quarters: [4]Quarter{
			NewQuarter(year, FirstQuarter, wd),
			NewQuarter(year, SecondQuarter, wd),
			NewQuarter(year, ThirdQuarter, wd),
			NewQuarter(year, FourthQuarter, wd),
		},
	}
}
