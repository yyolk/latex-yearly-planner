package mos

import (
	"bytes"
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/pages"
	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	types2 "github.com/kudrykv/latex-yearly-planner/app2/types"
	"github.com/kudrykv/latex-yearly-planner/app3/components"
	"github.com/kudrykv/latex-yearly-planner/app3/sections"
	"github.com/kudrykv/latex-yearly-planner/app3/types"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Planner struct {
	parameters Parameters
	year       calendar.Year
}

func New(parameters Parameters) *Planner {
	year := calendar.NewYear(parameters.Year, parameters.Weekday)

	return &Planner{
		parameters: parameters,
		year:       year,
	}
}

var ErrUnknownSection = fmt.Errorf("unknown section")

func (r *Planner) Generate() (types.NamedBuffers, error) {
	availableSections := r.sections()

	result := make(types.NamedBuffers, 0, len(r.parameters.Sections))

	for _, name := range r.parameters.Sections {
		section, ok := availableSections[name]
		if !ok {
			return nil, fmt.Errorf("lookup %s: %w", name, ErrUnknownSection)
		}

		buff, err := section()
		if err != nil {
			return nil, fmt.Errorf("run %s: %w", name, err)
		}

		result = append(result, types.NamedBuffer{Name: name, Buffer: buff})
	}

	return result, nil
}

func (r *Planner) sections() map[string]types2.SectionFunc {
	return map[string]types2.SectionFunc{
		common.QuarterliesSection:  r.quarterliesSection,
		common.MonthliesSection:    r.monthliesSection,
		common.WeekliesSection:     r.weekliesSection,
		common.DailiesSection:      r.dailiesSection,
		common.DailyNotesSection:   r.dailyNotesSection,
		common.DailyReflectSection: r.dailyReflectSection,
		common.NotesSection:        r.notesSection,
		common.ToDoSection:         r.todosSection,
	}
}

func (r *Planner) quarterliesSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	for _, quarter := range r.year.GetQuarters() {
		header, err := sections.NewMOSHeaderQuarterly(r.year, quarter, r.tabs(), r.parameters.MOSHeaderParameters)
		if err != nil {
			return nil, fmt.Errorf("new header: %w", err)
		}

		quarterly, err := sections.NewQuarterly(quarter, r.parameters.QuarterlyParameters)
		if err != nil {
			return nil, fmt.Errorf("new quarterly: %w", err)
		}

		if err = buffer.WriteBlocks(header, quarterly); err != nil {
			return nil, fmt.Errorf("write quarterly blocks: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *Planner) monthliesSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	var (
		header  sections.MOSHeaderDaily
		monthly sections.Monthly
		err     error
	)

	for _, month := range r.year.Months() {
		if header, err = sections.NewMOSHeaderMonthly(r.year, month, r.tabs(), r.parameters.MOSHeaderParameters); err != nil {
			return nil, fmt.Errorf("new header: %w", err)
		}

		if monthly, err = sections.NewMonthly(month, r.parameters.MonthlyParameters); err != nil {
			return nil, fmt.Errorf("new monthly: %w", err)
		}

		header = header.Title(monthly)

		if err = buffer.WriteBlocks(header, monthly); err != nil {
			return nil, fmt.Errorf("write monthly blocks: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *Planner) weekliesSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	var (
		header sections.MOSHeaderDaily
		err    error
	)

	for _, week := range r.year.Weeks() {
		if header, err = sections.NewMOSHeaderWeekly(r.year, week, r.tabs(), r.parameters.MOSHeaderParameters); err != nil {
			return nil, fmt.Errorf("new header: %w", err)
		}

		weekly, err := sections.NewWeekly(week, r.parameters.WeeklyParameters)
		if err != nil {
			return nil, fmt.Errorf("new weekly: %w", err)
		}

		header = header.Target(weekly)
		header = header.Title(weekly)

		if err = buffer.WriteBlocks(header, weekly); err != nil {
			return nil, fmt.Errorf("write weekly blocks: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *Planner) dailiesSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	var (
		daily   sections.Daily
		notes   sections.DailyNotes
		reflect sections.DailyReflect
		header  sections.MOSHeaderDaily
		err     error
	)

	for _, day := range r.year.Days() {
		if header, err = sections.NewMOSHeaderDaily(day, r.tabs(), r.parameters.MOSHeaderParameters); err != nil {
			return nil, fmt.Errorf("new header: %w", err)
		}

		if daily, err = sections.NewDaily(day, r.parameters.DailyParameters); err != nil {
			return nil, fmt.Errorf("new daily: %w", err)
		}

		header = header.Target(daily)
		header = header.Title(daily)

		if r.parameters.DailyNotesEnabled() {
			if notes, err = sections.NewDailyNotes(day, r.parameters.DailyNotesParameters); err != nil {
				return nil, fmt.Errorf("new daily notes: %w", err)
			}

			daily = daily.AppendNearNotesLine(fmt.Sprintf("$\\vert$ %s", notes.Link()))
		}

		if r.parameters.ReflectEnabled() {
			if reflect, err = sections.NewDailyReflect(day, r.parameters.DailyReflectParameters); err != nil {
				return nil, fmt.Errorf("new daily reflect: %w", err)
			}

			daily = daily.AppendNearNotesLine(fmt.Sprintf("\\hfill{}%s", reflect.Link()))
		}

		if err = buffer.WriteBlocks(header, daily); err != nil {
			return nil, fmt.Errorf("write daily blocks: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *Planner) dailyNotesSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	var (
		header sections.MOSHeaderDaily
		daily  sections.Daily
		notes  sections.DailyNotes
		err    error
	)

	for _, day := range r.year.Days() {
		if header, err = sections.NewMOSHeaderDaily(day, r.tabs(), r.parameters.MOSHeaderParameters); err != nil {
			return nil, fmt.Errorf("new header: %w", err)
		}

		if daily, err = sections.NewDaily(day, r.parameters.DailyParameters); err != nil {
			return nil, fmt.Errorf("new daily: %w", err)
		}

		if notes, err = sections.NewDailyNotes(day, r.parameters.DailyNotesParameters); err != nil {
			return nil, fmt.Errorf("new daily notes: %w", err)
		}

		header = header.Target(notes)
		header = header.LinkBack(daily)
		header = header.Repeat(notes)
		header = header.Title(daily)

		if err = buffer.WriteBlocks(header, notes); err != nil {
			return nil, fmt.Errorf("write daily notes blocks: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *Planner) dailyReflectSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	var (
		header  sections.MOSHeaderDaily
		reflect sections.DailyReflect
		daily   sections.Daily
		err     error
	)

	for _, day := range r.year.Days() {
		if header, err = sections.NewMOSHeaderDaily(day, r.tabs(), r.parameters.MOSHeaderParameters); err != nil {
			return nil, fmt.Errorf("new header: %w", err)
		}

		if daily, err = sections.NewDaily(day, r.parameters.DailyParameters); err != nil {
			return nil, fmt.Errorf("new daily: %w", err)
		}

		if reflect, err = sections.NewDailyReflect(day, r.parameters.DailyReflectParameters); err != nil {
			return nil, fmt.Errorf("new daily reflect: %w", err)
		}

		header = header.Target(reflect)
		header = header.LinkBack(daily)
		header = header.Title(daily)

		if err = buffer.WriteBlocks(header, reflect); err != nil {
			return nil, fmt.Errorf("write daily blocks: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *Planner) notesSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	header, err := sections.NewMOSHeaderIncomplete(r.year, r.tabs(targetNotes), r.parameters.MOSHeaderParameters)
	if err != nil {
		return nil, fmt.Errorf("new header: %w", err)
	}

	index, err := sections.NewIndex(r.parameters.NotesParameters.IndexParameters)
	if err != nil {
		return nil, fmt.Errorf("new index: %w", err)
	}

	notes, err := sections.NewNotes(index, r.parameters.NotesParameters)
	if err != nil {
		return nil, fmt.Errorf("new notes: %w", err)
	}

	index = index.ItemReferencePrefix("note-")

	for page := 1; page <= index.IndexPages(); page++ {
		index = index.CurrentPage(page)
		header = header.Title(index).Target(index)

		if err = buffer.WriteBlocks(header, index); err != nil {
			return nil, fmt.Errorf("write index blocks: %w", err)
		}
	}

	header, err = sections.NewMOSHeaderIncomplete(r.year, r.tabs(), r.parameters.MOSHeaderParameters)
	if err != nil {
		return nil, fmt.Errorf("new header: %w", err)
	}

	for page := 1; page <= index.ItemPages(); page++ {
		notes = notes.CurrentPage(page)
		header = header.Title(notes)
		header = header.Target(notes)
		header = header.LinkBack(index.IndexPageFromItemPage(page))

		if err = buffer.WriteBlocks(header, notes); err != nil {
			return nil, fmt.Errorf("write index blocks: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *Planner) todosSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	header, err := sections.NewMOSHeaderIncomplete(r.year, r.tabs(targetTodos), r.parameters.MOSHeaderParameters)
	if err != nil {
		return nil, fmt.Errorf("new header: %w", err)
	}

	index, err := sections.NewIndex(r.parameters.TodosParameters.IndexParameters)
	if err != nil {
		return nil, fmt.Errorf("new index: %w", err)
	}

	todos, err := sections.NewTodos(index, r.parameters.TodosParameters)
	if err != nil {
		return nil, fmt.Errorf("new todos: %w", err)
	}

	index = index.ItemReferencePrefix("todo-")

	for page := 1; page <= index.IndexPages(); page++ {
		index = index.CurrentPage(page)
		header = header.Title(index).Target(index)

		if err = buffer.WriteBlocks(header, index); err != nil {
			return nil, fmt.Errorf("write index blocks: %w", err)
		}
	}

	header, err = sections.NewMOSHeaderIncomplete(r.year, r.tabs(), r.parameters.MOSHeaderParameters)
	if err != nil {
		return nil, fmt.Errorf("new header: %w", err)
	}

	for page := 1; page <= index.ItemPages(); page++ {
		todos = todos.CurrentPage(page)
		header = header.Title(todos)
		header = header.Target(todos)
		header = header.LinkBack(index.IndexPageFromItemPage(page))

		if err = buffer.WriteBlocks(header, todos); err != nil {
			return nil, fmt.Errorf("write index blocks: %w", err)
		}
	}

	return buffer.Buffer, nil
}

const (
	targetNotes = iota + 1
	targetTodos
)

func (r *Planner) tabs(target ...int) components.Tabs {
	tabs := components.Tabs{{Text: "Calendar"}}

	if r.parameters.NotesEnabled() {
		focus := contains(target, targetNotes)

		tabs = append(tabs, components.Tab{Text: "Notes", Reference: "note-Index", Target: focus})
	}

	if r.parameters.TodosEnabled() {
		focus := contains(target, targetTodos)

		tabs = append(tabs, components.Tab{Text: "Todos", Reference: "todo-Index", Target: focus})
	}

	return tabs
}

func contains(list []int, target int) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}

	return false
}

func (r *Planner) RunTimes() int {
	return 2
}

func (r *Planner) Document() types.Document {
	return r.parameters.Document
}
