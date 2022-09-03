package sections

import "fmt"

type TodosParameters struct {
	IndexParameters IndexParameters
}

type Todos struct {
	parameters TodosParameters
	page       int
	index      Index
}

func NewTodos(index Index, parameters TodosParameters) (Todos, error) {
	return Todos{
		parameters: parameters,
		index:      index,
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
	return []string{"hello todo"}, nil
}
