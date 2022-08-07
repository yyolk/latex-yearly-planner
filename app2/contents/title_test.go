package contents_test

import (
	"testing"

	"github.com/kudrykv/latex-yearly-planner/app2/contents"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTitle(t *testing.T) {
	t.Parallel()

	Convey("should build a title", t, func() {
		title := contents.NewTitle("title")

		build, err := title.Build()
		So(err, ShouldBeNil)
		So(build, ShouldHaveLength, 1)
		So(build[0], ShouldEqual, expectedTitle)
	})
}

const expectedTitle = `\hspace{0pt}\vfil
\hfill\resizebox{.7\linewidth}{!}{title}%
\pagebreak`
