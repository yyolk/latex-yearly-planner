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
	yearStr      string
}

const (
	calendarText = "Calendar"
	toDoText     = "To Do"
	notesText    = "Notes"
)

func New(params common.Params) MonthsOnSides {
	return MonthsOnSides{
		year:    params.Year,
		weekday: params.Weekday,

		calendarYear: calendar.NewYear(params.Year, params.Weekday),
		yearStr:      strconv.Itoa(params.Year),
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
	buffer := pages.NewBuffer()

	if err := texsnippets.Execute(buffer, texsnippets.Title, map[string]string{"Title": r.yearStr}); err != nil {
		return nil, fmt.Errorf("execute template title: %w", err)
	}

	return buffer.Buffer, nil
}

func (r *MonthsOnSides) annualSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()
	header := r.headerWithTitleAndSelection(r.yearStr, calendarText)

	if err := buffer.WriteBlocks(header, annualContents{year: r.calendarYear}); err != nil {
		return nil, fmt.Errorf("write to buffer: %w", err)
	}

	return buffer.Buffer, nil
}

func (r *MonthsOnSides) quarterliesSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	for _, quarter := range r.calendarYear.Quarters {
		header := r.
			headerWithTitle(quarter.Name()).
			apply(headerSelectQuarter(quarter))

		if err := buffer.WriteBlocks(header, quarterlyContents{quarter: quarter}); err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *MonthsOnSides) monthliesSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	for _, month := range r.calendarYear.Months() {
		header := r.
			headerWithTitle(month.Month().String()).
			apply(headerSelectMonths(month.Month()))

		if err := buffer.WriteBlocks(header, monthlyContents{month: month}); err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *MonthsOnSides) weekliesSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	for _, week := range r.calendarYear.InWeeks() {
		header := r.
			headerWithTitle("Week " + strconv.Itoa(week.WeekNumber())).
			apply(headerSelectMonths(week.HeadMonth(), week.TailMonth()))

		if err := buffer.WriteBlocks(header, weeklyContents{week: week}); err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *MonthsOnSides) dailiesSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	for _, day := range r.calendarYear.Days() {
		header := r.
			headerWithTitle(day.Format("Mon Jan _2")).
			apply(headerSelectMonths(day.Month()))

		if err := buffer.WriteBlocks(header, dailyContents{day: day}); err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *MonthsOnSides) toDoSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()
	header := r.headerWithTitleAndSelection("To Do Index", toDoText)

	if err := buffer.WriteBlocks(header, todoIndex{}); err != nil {
		return nil, fmt.Errorf("write to buffer: %w", err)
	}

	for i := 1; i <= 100; i++ {
		header := r.headerWithTitle(strconv.Itoa(i))

		if err := buffer.WriteBlocks(header, todoContents{}); err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *MonthsOnSides) notesSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()
	header := r.headerWithTitleAndSelection("Notes Index", notesText)

	if err := buffer.WriteBlocks(header, notesIndex{}); err != nil {
		return nil, fmt.Errorf("write to buffer: %w", err)
	}

	for i := 1; i <= 100; i++ {
		header := r.headerWithTitle(strconv.Itoa(i))

		if err := buffer.WriteBlocks(header, notesContents{}); err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *MonthsOnSides) headerWithTitle(title string) header {
	return r.headerWithTitleAndSelection(title, "")
}

func (r *MonthsOnSides) headerWithTitleAndSelection(title string, selectText string) header {
	return r.baseHeader().apply(
		headerWithLeft(title),
		headerWithRight(r.rightCells().Ref(selectText).Slice()),
	)
}

func (r *MonthsOnSides) baseHeader() header {
	return newHeader(
		r.layout,
		r.ui,
		headerWithYear(r.calendarYear),
	)
}

func (r *MonthsOnSides) rightCells() cell.Cells {
	return cell.NewCells(calendarText, toDoText, notesText)
}
