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
	deviceName = "device-name"
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

		Commands: cli.Commands{
			&cli.Command{
				Name: "template",
				Subcommands: cli.Commands{
					&cli.Command{
						Name: "mos",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: deviceName, Required: true},
						},
						Action: func(appContext *cli.Context) error {
							device, err := devices.New(appContext.String(deviceName))
							if err != nil {
								return fmt.Errorf("new device: %w", err)
							}

							planner, err := planners.New(planners.NewParams(planners.MonthsOnSidesTemplate))
							if err != nil {
								return fmt.Errorf("new planner: %w", err)
							}

							if err = planner.GenerateFor(device); err != nil {
								return fmt.Errorf("generate: %w", err)
							}

							if err = planner.WriteTo("./out"); err != nil {
								return fmt.Errorf("write to ./out: %w", err)
							}

							if err = planner.Compile(appContext.Context); err != nil {
								return fmt.Errorf("compile: %w", err)
							}

							return nil
						},
					},
				},
			},
		},
	}

	return r
}

func (r *App) Run(args []string) error {
	return r.app.Run(args)
}
