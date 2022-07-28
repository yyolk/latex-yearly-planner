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

		args := []string{
			"./app", "template", "mos",
			//"--hand", "left",
			"--weekday", "1",
			"--show-frames",
			//"--show-links",
			"--device-name", "supernote_a5x",
			"--ui-path", "test.toml",
		}

		err := app2.New(in, out, errOut).Run(args)

		So(err, ShouldBeNil)
	})

}
