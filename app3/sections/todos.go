package sections

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app3/components"
)

type TodosParameters struct {
	IndexParameters IndexParameters
	TodosParameters components.TodosParameters
}

type Todos struct {
	parameters TodosParameters
	page       int
	index      Index
	todos      components.Todos
}

func NewTodos(index Index, parameters TodosParameters) (Todos, error) {
	todos, err := components.NewTodos(parameters.TodosParameters)
	if err != nil {
		return Todos{}, fmt.Errorf("new todos: %w", err)
	}

	return Todos{
		parameters: parameters,
		index:      index,
		todos:      todos,
	}, nil
}

func (r Todos) CurrentPage(page int) Todos {
	r.page = page

	return r
}

func (r Todos) Title() string {
	return fmt.Sprintf("Todos %d", r.page)
}

func (r Todos) Reference() string {
	return fmt.Sprintf("todo-%d", r.page)
}

func (r Todos) Build() ([]string, error) {
	return []string{r.todos.Build()}, nil
}
