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
		common.DailiesSection:      r.dailiesSection,
		common.DailyNotesSection:   r.dailyNotesSection,
		common.DailyReflectSection: r.dailyReflectSection,
		common.NotesSection:        r.notesSection,
	}
}

func (r *Planner) dailiesSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	tabs := components.Tabs{
		{Text: "Calendar"},
		{Text: "Notes"},
		{Text: "Todos"},
	}

	var (
		daily   sections.Daily
		notes   sections.DailyNotes
		reflect sections.DailyReflect
		header  sections.MOSHeaderDaily
		err     error
	)

	for _, day := range r.year.Days() {
		if header, err = sections.NewMOSHeaderDaily(day, tabs, r.parameters.MOSHeaderParameters); err != nil {
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

	tabs := components.Tabs{
		{Text: "Calendar"},
		{Text: "Notes"},
		{Text: "Todos"},
	}

	var (
		header sections.MOSHeaderDaily
		daily  sections.Daily
		notes  sections.DailyNotes
		err    error
	)

	for _, day := range r.year.Days() {
		if header, err = sections.NewMOSHeaderDaily(day, tabs, r.parameters.MOSHeaderParameters); err != nil {
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

	tabs := components.Tabs{
		{Text: "Calendar"},
		{Text: "Notes"},
		{Text: "Todos"},
	}

	var (
		header  sections.MOSHeaderDaily
		reflect sections.DailyReflect
		daily   sections.Daily
		err     error
	)

	for _, day := range r.year.Days() {
		if header, err = sections.NewMOSHeaderDaily(day, tabs, r.parameters.MOSHeaderParameters); err != nil {
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

	tabs := components.Tabs{
		{Text: "Calendar"},
		{Text: "Notes", Target: true},
		{Text: "Todos"},
	}

	header, err := sections.NewMOSHeaderIncomplete(r.year, tabs, r.parameters.MOSHeaderParameters)
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

	for page := 1; page <= index.IndexPages(); page++ {
		index = index.CurrentPage(page)
		header = header.Title(index)
		header = header.Target(index)

		if err = buffer.WriteBlocks(header, index); err != nil {
			return nil, fmt.Errorf("write index blocks: %w", err)
		}
	}

	tabs = components.Tabs{
		{Text: "Calendar"},
		{Text: "Notes"},
		{Text: "Todos"},
	}

	header, err = sections.NewMOSHeaderIncomplete(r.year, tabs, r.parameters.MOSHeaderParameters)
	if err != nil {
		return nil, fmt.Errorf("new header: %w", err)
	}

	for page := 1; page <= index.ItemPages(); page++ {
		notes = notes.CurrentPage(page)
		header = header.Title(notes)
		header = header.LinkBack(index.IndexPageFromItemPage(page))

		if err = buffer.WriteBlocks(header, notes); err != nil {
			return nil, fmt.Errorf("write index blocks: %w", err)
		}
	}

	return buffer.Buffer, nil
}

func (r *Planner) RunTimes() int {
	return 2
}

func (r *Planner) Document() types.Document {
	return r.parameters.Document
}
