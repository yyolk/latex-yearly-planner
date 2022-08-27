package components

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app3/types"
)

type ScheduleParameters struct {
	FromHour Hour
	ToHour   Hour
	Format   string

	LineHeight types.Millimeters
}

type Hour int

func (r Hour) Valid() bool {
	if r < 0 {
		return false
	}

	if r > 23 {
		return false
	}

	return true
}

var (
	ErrInvalidHour    = errors.New("invalid hour")
	ErrMisplacedHours = errors.New("to hour is less than from hour")
)

func (p ScheduleParameters) Test() error {
	if !p.FromHour.Valid() {
		return fmt.Errorf("from hour: %w", ErrInvalidHour)
	}

	if !p.ToHour.Valid() {
		return fmt.Errorf("to hour: %w", ErrInvalidHour)
	}

	if p.ToHour <= p.FromHour {
		return ErrMisplacedHours
	}

	return nil
}

type Schedule struct {
	parameters ScheduleParameters
	day        time.Time
}

func NewSchedule(day time.Time, parameters ScheduleParameters) (Schedule, error) {
	if err := parameters.Test(); err != nil {
		return Schedule{}, fmt.Errorf("test: %w", err)
	}

	return Schedule{
		day:        day,
		parameters: parameters,
	}, nil
}

func (r Schedule) Build() string {
	lineHeight := r.parameters.LineHeight
	pieces := make([]string, 0)

	for hour := r.parameters.FromHour; hour <= r.parameters.ToHour; hour++ {
		format := time.
			Date(r.day.Year(), r.day.Month(), r.day.Day(), int(hour), 0, 0, 0, time.UTC).
			Format(r.parameters.Format)

		pieces = append(pieces, fmt.Sprintf(scheduleHourFormat, lineHeight, format, lineHeight))
	}

	return strings.Join(pieces, "\n")
}

const scheduleHourFormat = `\parbox{0pt}{\vskip%s}%s\myLineLightGray` + "\n" + `\vskip%s\myLineGray`
