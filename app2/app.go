package app2

import (
	"os"

	"github.com/urfave/cli/v2"
)

func New() *cli.App {
	return &cli.App{
		Name: "plannergen",

		Reader:    os.Stdin,
		Writer:    os.Stdout,
		ErrWriter: os.Stderr,

		Action: mainAction,
	}
}

func mainAction(appContext *cli.Context) error {
	// dimensions:
	// - device (supernote, remarkable (vanilla, ddvk), ?boox)
	//     this defines dimensions and internal layout boundaries: leave empty space for control elements
	// - template - which template to use
	//     sub-dependent are enabled template sections
	//     based on what is selected, some links should or should not be displayed, etc

	// get / parse configs
	// create tex files
	// run latex

	return nil
}
