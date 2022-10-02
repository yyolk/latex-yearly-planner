package main

import (
	"strings"
	"testing"

	"github.com/kudrykv/latex-yearly-planner/app"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMoS3(t *testing.T) {
	Convey("mos", t, func() {
		in := strings.NewReader("")
		out := &strings.Builder{}
		errOut := &strings.Builder{}

		args := []string{
			"./app", "template", "mos",
			"--layout-path", "cfg/sn_a5x_mos.toml",
			"--parameters-path", "cfg/sn_a5x_mos.toml",
		}

		err := app.New(in, out, errOut).Run(args)

		So(err, ShouldBeNil)
	})
}
