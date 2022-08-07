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

func TestMoS2(t *testing.T) {
	Convey("mos2", t, func() {
		in := strings.NewReader("")
		out := &strings.Builder{}
		errOut := &strings.Builder{}

		args := []string{
			"./app", "template", "mos2",
			//"--hand", "left",
			"--weekday", "1",
			"--show-frames",
			//"--show-links",
			"--device-name", "supernote_a5x",
		}

		err := app2.New(in, out, errOut).Run(args)

		So(err, ShouldBeNil)
	})
}

func TestBreadcrumb(t *testing.T) {
	Convey("Breadcrumb", t, func() {
		in := strings.NewReader("")
		out := &strings.Builder{}
		errOut := &strings.Builder{}

		args := []string{
			"./app", "template", "breadcrumb",
			//"--hand", "left",
			"--weekday", "1",
			//"--show-frames",
			//"--show-links",
			"--device-name", "supernote_a5x",
		}

		err := app2.New(in, out, errOut).Run(args)

		So(err, ShouldBeNil)
	})
}
