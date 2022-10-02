package calendar

import (
	"strconv"
	"time"
)

type Quarters []Quarter

type Quarter struct {
	Months [3]Month
	number int
	year   *Year
}

func newQuarter(year *Year, quarter int, weekday time.Weekday) Quarter {
	if quarter < FirstQuarter || quarter > FourthQuarter {
		return Quarter{} //nolint:exhaustivestruct
	}

	monthNumber := time.Month((quarter+1)*3 - 2) //nolint:gomnd

	calendarQuarter := Quarter{
		year:   year,
		number: quarter + 1,
	}

	for month := monthNumber; month <= monthNumber+2; month++ {
		calendarQuarter.Months[month-monthNumber] = newMonth(year, &calendarQuarter, month, weekday)
	}

	return calendarQuarter
}

func NewQuarter(year int, qrtr int, wd time.Weekday) Quarter {
	if qrtr < FirstQuarter || qrtr > FourthQuarter {
		return Quarter{} //nolint:exhaustivestruct
	}

	mo := time.Month((qrtr+1)*3 - 2) //nolint:gomnd

	return Quarter{
		number: qrtr + 1,

		Months: [3]Month{
			NewMonth(year, mo, wd),
			NewMonth(year, mo+1, wd),
			NewMonth(year, mo+2, wd), //nolint:gomnd
		},
	}
}

func (r Quarter) Name() string {
	return "Q" + strconv.Itoa(r.number)
}

func (r Quarter) Number() int {
	return r.number
}

func (r Quarter) Year() Year {
	return *r.year
}
