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
		common.DailiesSection:    r.dailiesSection,
		common.DailyNotesSection: r.dailyNotesSection,
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
		daily  sections.Daily
		header sections.MOSHeaderDaily
		err    error
	)

	for _, day := range r.year.Days() {
		if header, err = sections.NewMOSHeaderDaily(day, tabs, r.parameters.MOSHeaderParameters); err != nil {
			return nil, fmt.Errorf("new header: %w", err)
		}

		if daily, err = sections.NewDaily(day, r.parameters.DailyParameters); err != nil {
			return nil, fmt.Errorf("new daily: %w", err)
		}

		header = header.Target(daily)

		if r.parameters.DailyNotesEnabled() {
			notes := sections.NewDailyNotes(day, r.parameters.DailyNotesParameters)

			daily = daily.NearNotesLine(fmt.Sprintf("— %s", notes.Link("More")))
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
		err    error
	)

	for _, day := range r.year.Days() {
		if header, err = sections.NewMOSHeaderDaily(day, tabs, r.parameters.MOSHeaderParameters); err != nil {
			return nil, fmt.Errorf("new header: %w", err)
		}

		if daily, err = sections.NewDaily(day, r.parameters.DailyParameters); err != nil {
			return nil, fmt.Errorf("new daily: %w", err)
		}

		notes := sections.NewDailyNotes(day, r.parameters.DailyNotesParameters)
		header = header.Target(notes)
		header = header.LinkBack(daily)
		header = header.Repeat(notes)

		if err = buffer.WriteBlocks(header, notes); err != nil {
			return nil, fmt.Errorf("write daily notes blocks: %w", err)
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
