package components_test

import (
	"testing"

	"github.com/kudrykv/latex-yearly-planner/app/components"
	. "github.com/kudrykv/latex-yearly-planner/app/test"
	"github.com/stretchr/testify/assert"
)

func TestTabLine_Build(t *testing.T) {
	t.Parallel()

	tabs := components.Tabs{
		{Text: "one", Reference: "ref-one", Target: true},
		{Text: "two", Reference: "ref-two", Target: false},
	}
	parameters := components.TabLineParameters{
		VerticalSpacing:     5,
		SpaceBetweenColumns: 3,
	}

	tabLine := components.NewTabLine(tabs, parameters)

	assert.Equal(t, Fixture("tab_line"), tabLine.Build())
}
