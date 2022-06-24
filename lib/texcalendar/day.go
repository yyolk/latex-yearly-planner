package texcalendar

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/tex/ref"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Days []Day

func (d Days) BuildLittle() []string {
	out := make([]string, 0, len(d))

	for _, day := range d {
		if day.Day.IsZero() {
			out = append(out, "")

			continue
		}

		name := day.Day.Format("_2")
		out = append(out, ref.NewLinkWithRef(name, day.Ref()).Build())
	}

	return out
}

func (d Days) BuildLarge() []string {
	out := make([]string, 0, len(d))

	for _, day := range d {
		if day.Day.IsZero() {
			out = append(out, "")

			continue
		}

		name := `{\renewcommand{\arraystretch}{1.2}\begin{tabular}{@{}p{5mm}@{}|}\hfil{}` + day.Day.Format("_2") + `\\ \hline\end{tabular}}`
		out = append(out, ref.NewLinkWithRef(name, day.Ref()).Build())
	}

	return out
}

type Day struct {
	Day        calendar.Day
	parameters Parameters
}

func NewDay(day calendar.Day, options ...ApplyToParameters) Day {
	parameters := Parameters{}

	for _, option := range options {
		option(&parameters)
	}

	return Day{Day: day, parameters: parameters}
}

func (d Day) NameAndDate() string {
	return d.Day.Format("Monday, _2")
}

func (d Day) Ref() string {
	return d.Day.Format(time.RFC3339)
}

func (d Day) Month() time.Month {
	return d.Day.Month()
}

func (d Day) Week() Week {
	return NewWeek(*d.Day.Week())
}
