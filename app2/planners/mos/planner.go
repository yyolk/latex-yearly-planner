package mos

import (
	"bytes"
	"fmt"
	"strconv"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
	"github.com/kudrykv/latex-yearly-planner/app2/pages"
	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/app2/tex/cell"
	"github.com/kudrykv/latex-yearly-planner/app2/texsnippets"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type MonthsOnSides struct {
	device devices.Device
	layout common.Layout

	ui           mosUI
	year         int
	weekday      time.Weekday
	calendarYear calendar.Year
}

const (
	calendarCell = "Calendar"
	toDoCell     = "To Do"
	notesCell    = "Notes"
)

func New(params common.Params) MonthsOnSides {
	return MonthsOnSides{
		year:    params.Year,
		weekday: params.Weekday,

		calendarYear: calendar.NewYear(params.Year, params.Weekday),
	}
}

func (r *MonthsOnSides) SetLayout(layout common.Layout) {
	r.layout = layout
}

func (r MonthsOnSides) Layout() common.Layout {
	return r.layout
}

func (r *MonthsOnSides) PrepareDetails(device devices.Device) error {
	r.device = device

	switch device.(type) {
	case *devices.SupernoteA5X:
		r.ui = mosUI{
			HeaderMarginNotesArrayStretch:  "2.042",
			HeaderMarginNotesMonthsWidth:   "15.7cm",
			HeaderMarginNotesQuartersWidth: "5cm",
			HeaderArrayStretch:             "1.8185",
		}
	default:
		return fmt.Errorf("%T: %w", device, common.UnknownDeviceTypeErr)
	}

	return nil
}

type sectionFunc func() (*bytes.Buffer, error)

func (r MonthsOnSides) Sections() map[string]sectionFunc {
	return map[string]sectionFunc{
		common.TitleSection:       r.titleSection,
		common.AnnualSection:      r.annualSection,
		common.QuarterliesSection: r.quarterliesSection,
		common.MonthliesSection:   r.monthliesSection,
		common.WeekliesSection:    r.weekliesSection,
		common.DailiesSection:     r.dailiesSection,
		common.ToDoSection:        r.toDoSection,
		common.NotesSection:       r.notesSection,
	}
}

func (r *MonthsOnSides) titleSection() (*bytes.Buffer, error) {
	buffer := &bytes.Buffer{}

	title := strconv.Itoa(r.year)
	if err := texsnippets.Execute(buffer, texsnippets.Title, map[string]string{"Title": title}); err != nil {
		return nil, fmt.Errorf("execute template title: %w", err)
	}

	return buffer, nil
}

func (r *MonthsOnSides) annualSection() (*bytes.Buffer, error) {
	header := newHeader(
		r.layout,
		r.ui,
		headerWithYear(r.calendarYear),
		headerWithLeft(strconv.Itoa(r.year)),
		headerWithRight(r.headerRightCells().Ref(calendarCell).Slice()),
	)

	buffer, err := writeToBuffer(&bytes.Buffer{}, header, annualContents{year: r.calendarYear})
	if err != nil {
		return nil, fmt.Errorf("write to buffer: %w", err)
	}

	return buffer, nil
}

func (r *MonthsOnSides) quarterliesSection() (*bytes.Buffer, error) {
	var (
		buffer = &bytes.Buffer{}
		err    error
	)

	for _, quarter := range r.calendarYear.Quarters {
		header := newHeader(
			r.layout,
			r.ui,
			headerWithYear(r.calendarYear),
			headerWithLeft(quarter.Name()),
			headerWithRight(r.headerRightCells().Slice()),
			headerSelectQuarter(quarter),
		)

		if buffer, err = writeToBuffer(buffer, header, quarterlyContents{quarter: quarter}); err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer, nil
}

func (r *MonthsOnSides) monthliesSection() (*bytes.Buffer, error) {
	var (
		buffer = &bytes.Buffer{}
		err    error
	)

	for _, month := range r.calendarYear.Months() {
		header := newHeader(
			r.layout,
			r.ui,
			headerWithYear(r.calendarYear),
			headerWithLeft(month.Month().String()),
			headerWithRight(r.headerRightCells().Slice()),
			headerSelectMonths(month.Month()),
		)

		if buffer, err = writeToBuffer(buffer, header, monthlyContents{month: month}); err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer, nil
}

func (r *MonthsOnSides) weekliesSection() (*bytes.Buffer, error) {
	var (
		buffer = &bytes.Buffer{}
		err    error
	)

	weeks := r.calendarYear.InWeeks()

	for _, week := range weeks {
		header := newHeader(
			r.layout,
			r.ui,
			headerWithYear(r.calendarYear),
			headerWithLeft("Week "+strconv.Itoa(week.WeekNumber())),
			headerWithRight(r.headerRightCells().Slice()),
			headerSelectMonths(week.HeadMonth(), week.TailMonth()),
		)

		if buffer, err = writeToBuffer(buffer, header, weeklyContents{week: week}); err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer, nil
}

func (r *MonthsOnSides) dailiesSection() (*bytes.Buffer, error) {
	var (
		buffer = &bytes.Buffer{}
		err    error
	)

	for _, day := range r.calendarYear.Days() {
		header := newHeader(
			r.layout,
			r.ui,
			headerWithYear(r.calendarYear),
			headerWithLeft(day.Format("Mon Jan _2")),
			headerWithRight(r.headerRightCells().Slice()),
			headerSelectMonths(day.Month()),
		)

		if buffer, err = writeToBuffer(buffer, header, dailyContents{day: day}); err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer, nil
}

func (r *MonthsOnSides) toDoSection() (*bytes.Buffer, error) {
	var (
		buffer = &bytes.Buffer{}
		err    error
	)

	header := newHeader(
		r.layout,
		r.ui,
		headerWithYear(r.calendarYear),
		headerWithLeft("To Do Index"),
		headerWithRight(r.headerRightCells().Ref(toDoCell).Slice()),
	)

	if buffer, err = writeToBuffer(buffer, header, todoIndex{}); err != nil {
		return nil, fmt.Errorf("write to buffer: %w", err)
	}

	for i := 1; i <= 100; i++ {
		header := newHeader(
			r.layout,
			r.ui,
			headerWithYear(r.calendarYear),
			headerWithLeft(strconv.Itoa(i)),
			headerWithRight(r.headerRightCells().Slice()),
		)

		if buffer, err = writeToBuffer(buffer, header, todoContents{}); err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer, nil
}

func (r *MonthsOnSides) notesSection() (*bytes.Buffer, error) {
	var (
		buffer = &bytes.Buffer{}
		err    error
	)

	header := newHeader(
		r.layout,
		r.ui,
		headerWithYear(r.calendarYear),
		headerWithLeft("Notes Index"),
		headerWithRight(r.headerRightCells().Ref(notesCell).Slice()),
	)

	if buffer, err = writeToBuffer(buffer, header, notesIndex{}); err != nil {
		return nil, fmt.Errorf("write to buffer: %w", err)
	}

	for i := 1; i <= 100; i++ {
		header := newHeader(
			r.layout,
			r.ui,
			headerWithYear(r.calendarYear),
			headerWithLeft(strconv.Itoa(i)),
			headerWithRight(r.headerRightCells().Slice()),
		)

		if buffer, err = writeToBuffer(buffer, header, notesContents{}); err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer, nil
}

func (r *MonthsOnSides) headerRightCells() cell.Cells {
	return cell.NewCells(calendarCell, toDoCell, notesCell)
}

func writeToBuffer(buffer *bytes.Buffer, blocks ...pages.Block) (*bytes.Buffer, error) {
	compiledPages, err := pages.NewPage(blocks...).Build()
	if err != nil {
		return nil, fmt.Errorf("build new page: %w", err)
	}

	for _, page := range compiledPages {
		buffer.WriteString(page + "\n\n" + `\pagebreak{}` + "\n")
	}

	return buffer, nil
}
