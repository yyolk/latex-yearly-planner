package planners

import (
	"errors"
	"fmt"
)

const (
	MonthsOnSidesTemplate = "mos"
)

var UnknownTemplateName = errors.New("unknown planner name")

func New(params Params) (*Planner, error) {
	var builder MonthsOnSides

	switch params.Name {
	case MonthsOnSidesTemplate:
		builder = newMonthOnSides(params.TemplateData)
	default:
		return nil, fmt.Errorf("%s: %w", params.Name, UnknownTemplateName)
	}

	return &Planner{
		params:      params,
		futureFiles: nil,
		dir:         "",
		builder:     builder,
	}, nil
}
