package app2

import (
	"fmt"
	"io"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
	"github.com/kudrykv/latex-yearly-planner/app2/planners"
	"github.com/urfave/cli/v2"
)

type App struct {
	app *cli.App
}

const (
	yearFlag     = "year"
	sectionsFlag = "sections"

	deviceNameFlag = "device-name"

	handFlag = "hand"
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
							&cli.IntFlag{Name: yearFlag, Value: time.Now().Year()},
							&cli.StringFlag{Name: deviceNameFlag, Required: true},
							&cli.StringFlag{Name: handFlag, Value: "right"},
							&cli.StringSliceFlag{
								Name: sectionsFlag,
								Value: cli.NewStringSlice(
									//planners.TitleSection,
									planners.AnnualSection,
									planners.QuarterliesSection,
									//planners.MonthliesSection,
									//planners.WeekliesSection,
									//planners.DailiesSection,
								),
							},
						},
						Action: func(appContext *cli.Context) error {
							device, err := devices.New(appContext.String(deviceNameFlag))
							if err != nil {
								return fmt.Errorf("new device: %w", err)
							}

							hand := planners.RightHand
							if appContext.String(handFlag) == "left" {
								hand = planners.LeftHand
							}

							params := planners.NewParams(planners.MonthsOnSidesTemplate)
							params.TemplateData.Apply(
								planners.WithYear(appContext.Int(yearFlag)),
								planners.WithSections(appContext.StringSlice(sectionsFlag)),
							)

							planner, err := planners.New(params)
							if err != nil {
								return fmt.Errorf("new planner: %w", err)
							}

							if err = planner.GenerateFor(device, hand); err != nil {
								return fmt.Errorf("generate: %w", err)
							}

							if err = planner.WriteTeXTo("./out"); err != nil {
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
