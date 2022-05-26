package main

import (
	"strings"
	"testing"

	"github.com/kudrykv/latex-yearly-planner/app2"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDefault(t *testing.T) {
	Convey("Default", t, func() {
		in := strings.NewReader("")
		out := &strings.Builder{}
		errOut := &strings.Builder{}

		err := app2.New(in, out, errOut).Run([]string{"./app", "--device-name", "supernote_a5x"})

		So(err, ShouldBeNil)
	})

}
