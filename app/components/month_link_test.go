package components_test

import (
	"testing"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app/components"
	"github.com/stretchr/testify/assert"
)

func TestMonthLink_String(t *testing.T) {
	t.Parallel()

	table := map[string]struct {
		link components.MonthLink
		want string
	}{
		"valid month": {
			link: components.NewMonthLink(time.January),
			want: `\hyperlink{January}{January}`,
		},

		"invalid month": {
			link: components.NewMonthLink(time.Month(13)),
			want: `\hyperlink{%!Month(13)}{%!Month(13)}`,
		},
	}

	for name, test := range table {
		name, test := name, test

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.want, test.link.String())
		})
	}
}
