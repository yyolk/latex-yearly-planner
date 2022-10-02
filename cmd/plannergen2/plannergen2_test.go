package main

import (
	"strings"
	"testing"

	"github.com/kudrykv/latex-yearly-planner/app3"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMoS3(t *testing.T) {
	Convey("mos", t, func() {
		in := strings.NewReader("")
		out := &strings.Builder{}
		errOut := &strings.Builder{}

		args := []string{
			"./app", "template", "mos2",
			"--layout-path", "cfg2/sn_a5x_mos.toml",
			"--parameters-path", "cfg2/sn_a5x_mos.toml",
		}

		err := app3.New(in, out, errOut).Run(args)

		So(err, ShouldBeNil)
	})
}
