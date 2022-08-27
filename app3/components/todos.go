package components

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app3/types"
)

type TodosParameters struct {
	Number int

	LineHeight types.Millimeters
}

var ErrNegativeTodos = errors.New("number of todos must be positive")

func (r TodosParameters) Test() error {
	if r.Number < 0 {
		return ErrNegativeTodos
	}

	return nil
}

type Todos struct {
	parameters TodosParameters
}

func NewTodos(parameters TodosParameters) (Todos, error) {
	if err := parameters.Test(); err != nil {
		return Todos{}, fmt.Errorf("test: %w", err)
	}

	return Todos{parameters}, nil
}

func (r Todos) Build() string {
	pieces := make([]string, 0, r.parameters.Number)

	for i := 1; i <= r.parameters.Number; i++ {
		pieces = append(pieces, fmt.Sprintf(`\parbox{0pt}{\vskip%s}$\square$\myLineGray`, r.parameters.LineHeight))
	}

	return strings.Join(pieces, "\n")
}
