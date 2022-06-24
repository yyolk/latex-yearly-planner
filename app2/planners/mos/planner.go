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
	"github.com/kudrykv/latex-yearly-planner/app2/tex/ref"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type MonthsOnSides struct {
	device devices.Device
	layout common.Layout

	ui      ui
	weekday time.Weekday
	year    texcalendar.Year
}

const (
	calendarText = "Calendar"
	toDoText     = "To Do"
	notesText    = "Notes"
)

func New(params common.Params) MonthsOnSides {
	return MonthsOnSides{
		weekday: params.Weekday,

		year: texcalendar.NewYear(params.Year, texcalendar.WithParameters(texcalendar.Parameters{
			Hand:       params.Hand,
			Weekday:    params.Weekday,
			FirstMonth: time.January,
		})),
	}
}

func (r MonthsOnSides) Layout() common.Layout {
	return r.layout
}

func (r *MonthsOnSides) PrepareDetails(device devices.Device, layout common.Layout) error {
	r.layout = layout
	r.device = device

	var err error
	if r.ui, err = newUI(device); err != nil {
		return fmt.Errorf("new ui: %w", err)
	}

	return nil
}

type sectionFunc func() (*bytes.Buffer, error)

func (r MonthsOnSides) Sections() map[string]sectionFunc {
	return map[string]sectionFunc{
		common.TitleSection:        r.titleSection,
		common.AnnualSection:       r.annualSection,
		common.QuarterliesSection:  r.quarterliesSection,
		common.MonthliesSection:    r.monthliesSection,
		common.WeekliesSection:     r.weekliesSection,
		common.DailiesSection:      r.dailiesSection,
		common.DailyNotesSection:   r.dailyNotesSection,
		common.DailyReflectSection: r.dailyReflectSection,
		common.ToDoSection:         r.toDoSection,
		common.NotesSection:        r.notesSection,
	}
}

func (r *MonthsOnSides) titleSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	if err := buffer.WriteBlocks(&titleContents{title: r.year.Name()}); err != nil {
		return nil, fmt.Errorf("write to buffer: %w", err)
	}

	return buffer.Buffer, nil
}

func (r *MonthsOnSides) annualSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()
	header := r.headerWithTitleAndSelection(r.year.Name(), calendarText)

	if err := buffer.WriteBlocks(header, annualContents{hand: r.layout.Hand, year: r.year}); err != nil {
		return nil, fmt.Errorf("write to buffer: %w", err)
	}

	return buffer.Buffer, nil
}

func (r *MonthsOnSides) quarterliesSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	for _, quarter := range r.year.Quarters() {
		header := r.
			headerWithTitle(quarter.Name()).
			apply(headerSelectQuarter(quarter))

		if err := buffer.WriteBlocks(header, quarterlyContents{hand: r.layout.Hand, quarter: quarter}); err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *MonthsOnSides) monthliesSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	for _, month := range r.year.Months() {
		header := r.
			headerWithTitle(month.Month().String()).
			apply(headerSelectMonths(month.Month()))

		if err := buffer.WriteBlocks(header, monthlyContents{month: month, hand: r.layout.Hand}); err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *MonthsOnSides) weekliesSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	for _, week := range r.year.InWeeks() {
		header := r.
			headerWithTitle(ref.NewTargetWithRef(week.Title(), week.Ref()).Build()).
			apply(headerSelectMonths(r.highlightedMonths(week)...))

		if err := buffer.WriteBlocks(header, weeklyContents{week: week}); err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *MonthsOnSides) highlightedMonths(week texcalendar.Week) []time.Month {
	switch {
	case week.First():
		return []time.Month{week.TailMonth()}
	case week.Last():
		return []time.Month{week.HeadMonth()}
	default:
		return []time.Month{week.HeadMonth(), week.TailMonth()}
	}
}

func (r *MonthsOnSides) dailiesSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	for _, day := range r.year.Days() {
		week := day.Week()
		weekCell := cell.New(week.Title()).RefAs(week.Ref())

		header := r.headerWithTitle(ref.NewTargetWithRef(day.NameAndDate(), day.Ref()).Build()).apply(
			headerSelectMonths(day.Month()),
			headerAddAction(weekCell),
		)

		if err := buffer.WriteBlocks(header, dailyContents{hand: r.layout.Hand, day: day}); err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *MonthsOnSides) dailyNotesSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	for _, day := range r.year.Days() {
		week := day.Week()
		weekCell := cell.New(week.Title()).RefAs(week.Ref())

		header := r.
			headerWithTitle(ref.NewLinkWithRef(day.NameAndDate(), day.Ref()).Build()+ref.NewTargetWithRef("", day.Ref()+"-notes").Build()).
			apply(
				headerSelectMonths(day.Month()),
				headerAddAction(weekCell),
			)

		if err := buffer.WriteBlocks(header, dailyNotesContents{day: day}); err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *MonthsOnSides) dailyReflectSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	for _, day := range r.year.Days() {
		week := day.Week()
		weekCell := cell.New(week.Title()).RefAs(week.Ref())

		header := r.
			headerWithTitle(ref.NewLinkWithRef(day.NameAndDate(), day.Ref()).Build()+ref.NewTargetWithRef("", day.Ref()+"-reflect").Build()).
			apply(
				headerSelectMonths(day.Month()),
				headerAddAction(weekCell),
			)

		if err := buffer.WriteBlocks(header, dailyReflectContents{day: day}); err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *MonthsOnSides) toDoSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()
	header := r.headerWithTitleAndSelection("To Do Index", toDoText)

	if err := buffer.WriteBlocks(header.repeat(2), index{refPrefix: "todo"}); err != nil {
		return nil, fmt.Errorf("write to buffer: %w", err)
	}

	for i := 1; i <= 115; i++ {
		header := r.headerWithTitle(ref.NewTargetWithRef("To Do "+strconv.Itoa(i), "todo-"+strconv.Itoa(i)).Build())

		if err := buffer.WriteBlocks(header, todoContents{}); err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *MonthsOnSides) notesSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()
	header := r.headerWithTitleAndSelection("Notes Index", notesText)

	if err := buffer.WriteBlocks(header.repeat(2), index{refPrefix: "notes"}); err != nil {
		return nil, fmt.Errorf("write to buffer: %w", err)
	}

	for i := 1; i <= 115; i++ {
		header := r.headerWithTitle(ref.NewTargetWithRef("Notes "+strconv.Itoa(i), "notes-"+strconv.Itoa(i)).Build())

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
		headerWithTitle(title),
		headerWithActions(r.rightCells().Ref(selectText)),
	)
}

func (r *MonthsOnSides) baseHeader() header {
	return newHeader(
		r.layout,
		r.ui,
		headerWithTexYear(r.year),
		headerWithHand(r.layout.Hand),
	)
}

func (r *MonthsOnSides) rightCells() cell.Cells {
	return cell.NewCells(calendarText, toDoText, notesText)
}

func (r MonthsOnSides) RunTimes() int {
	// MoS need to be compiled two times to correctly position margin notes
	return 2
}
