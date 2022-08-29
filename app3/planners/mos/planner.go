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
		common.DailiesSection: r.dailiesSection,
	}
}

func (r *Planner) dailiesSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	var (
		daily sections.Daily
		err   error
	)

	for _, day := range r.year.Days() {
		tabs := components.Tabs{
			{Text: "Calendar"},
			{Text: "Notes"},
			{Text: "Todos"},
		}
		header := sections.NewMOSHeaderDaily(day, tabs, r.parameters.MOSHeaderParameters)

		if daily, err = sections.NewDaily(day, r.parameters.DailyParameters); err != nil {
			return nil, fmt.Errorf("new daily: %w", err)
		}

		if err = buffer.WriteBlocks(header, daily); err != nil {
			return nil, fmt.Errorf("write blocks: %w", err)
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
