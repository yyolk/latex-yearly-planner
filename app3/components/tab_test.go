package components_test

import (
	"testing"

	"github.com/kudrykv/latex-yearly-planner/app3/components"
	"github.com/stretchr/testify/assert"
)

func TestTab_String(t *testing.T) {
	tests := map[string]struct {
		tab  components.Tab
		want string
	}{
		"name": {
			tab:  components.Tab{Text: "hello"},
			want: `\hyperlink{hello}{hello}`,
		},
		"reference": {
			tab:  components.Tab{Text: "hello", Reference: "world"},
			want: `\hyperlink{world}{hello}`,
		},
		"target": {
			tab:  components.Tab{Text: "hello", Target: true},
			want: `\cellcolor{black}{\textcolor{white}{hello}}`,
		},
		"marked as target has no, uh, target": {
			tab:  components.Tab{Text: "hello", Target: true, Reference: "world"},
			want: `\cellcolor{black}{\textcolor{white}{hello}}`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.want, test.tab.String())
		})
	}
}
