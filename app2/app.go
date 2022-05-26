package app2

import (
	"fmt"
	"io"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
	"github.com/kudrykv/latex-yearly-planner/app2/planners"
	"github.com/urfave/cli/v2"
)

type App struct {
	app *cli.App
}

const (
	deviceName  = "device-name"
	plannerName = "planner-name"
)

func New(reader io.Reader, writer, errWriter io.Writer) *App {
	return (&App{}).
		setupCli(reader, writer, errWriter)
}

func (r *App) setupCli(reader io.Reader, writer, errWriter io.Writer) *App {
	r.app = &cli.App{
		Name: "plannergen",

		Reader:    reader,
		Writer:    writer,
		ErrWriter: errWriter,
		Flags:     r.flags(),

		Action: r.mainAction,
	}

	return r
}

func (r *App) flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{Name: deviceName, Required: true},
		&cli.StringFlag{Name: plannerName, Required: true},
	}
}

func (r *App) mainAction(appContext *cli.Context) error {
	device, err := devices.New(appContext.String(deviceName))
	if err != nil {
		return fmt.Errorf("new device: %w", err)
	}

	template, err := planners.New(appContext.String(plannerName))
	if err != nil {
		return fmt.Errorf("new template: %w", err)
	}

	if err = template.GenerateFor(device); err != nil {
		return fmt.Errorf("generate: %w", err)
	}

	return nil
}

func (r *App) Run(args []string) error {
	return r.app.Run(args)
}
