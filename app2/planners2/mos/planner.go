package mos

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/types"
)

type Planner struct {
	layout     types.Layout
	parameters Parameters
}

func New(layout types.Layout) (*Planner, error) {
	parameters, ok := layout.Misc.(Parameters)
	if !ok {
		return nil, fmt.Errorf("expected Parameters, got %T", layout.Misc)
	}

	planner := &Planner{
		layout:     layout,
		parameters: parameters,
	}

	for section := range planner.sections() {
		if !Contains(parameters.enabledSections, section) {
			return nil, fmt.Errorf("unknown section %s", section)
		}
	}

	return planner, nil
}

func (r *Planner) BuildData() (types.NamedDatas, error) {
	sections := r.sections()

	result := make(types.NamedDatas, 0, len(r.parameters.enabledSections))

	for _, name := range r.parameters.enabledSections {
		section, ok := sections[name]
		if !ok {
			panic(fmt.Sprintf("unknown section %s", name))
		}

		buff, err := section()
		if err != nil {
			return nil, fmt.Errorf("%s: %w", name, err)
		}

		result = append(result, types.NamedData{Name: name, Data: buff.Bytes()})
	}

	return result, nil
}

func (r *Planner) sections() map[string]types.SectionFunc {
	panic("here will be sections")
}

func (r *Planner) RunTimes() int {
	return 2
}

func Contains[T comparable](slice []T, item T) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}

	return false
}
