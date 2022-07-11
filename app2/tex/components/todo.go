package components

import (
	"fmt"
	"strings"
)

type Todos []Todo

func (r Todos) Build() string {
	todos := make([]string, 0, len(r))

	for _, todo := range r {
		todos = append(todos, todo.Build())
	}

	return strings.Join(todos, "\n")
}

func NewTodos(number int) Todos {
	todos := make(Todos, 0, number)

	for i := 0; i < number; i++ {
		todos = append(todos, NewTodo())
	}

	return todos
}

type Todo struct {
	Height string
}

func NewTodo() Todo {
	return Todo{}
}

func (r Todo) Build() string {
	r.prepare()

	return fmt.Sprintf(`\parbox{0pt}{\vskip%s}$\square$\myLineGray`, r.Height)
}

func (r *Todo) prepare() {
	if len(r.Height) == 0 {
		r.Height = "5mm"
	}
}
