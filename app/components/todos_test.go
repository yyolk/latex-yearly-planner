package components_test

import (
	"testing"

	"github.com/kudrykv/latex-yearly-planner/app/components"
	"github.com/kudrykv/latex-yearly-planner/app/test"
	"github.com/stretchr/testify/assert"
)

func TestTodos_Build(t *testing.T) {
	t.Parallel()

	parameters := components.TodosParameters{
		Number:     3,
		LineHeight: 5,
	}

	todos, err := components.NewTodos(parameters)

	assert.NoError(t, err)
	assert.Equal(t, test.Fixture("todos"), todos.Build())
}
