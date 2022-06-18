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

		args := []string{"./app", "template", "mos", "--hand", "right", "--weekday", "1", "--device-name", "supernote_a5x"}

		err := app2.New(in, out, errOut).Run(args)

		So(err, ShouldBeNil)
	})

}
