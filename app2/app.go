package app2

import (
	"fmt"
	"io"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
	"github.com/kudrykv/latex-yearly-planner/app2/planners"
	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/urfave/cli/v2"
)

type App struct {
	app *cli.App
}

const (
	yearFlag     = "year"
	sectionsFlag = "sections"
	weekdayFlag  = "weekday"

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
							&cli.IntFlag{Name: weekdayFlag, Value: 0},
							&cli.StringSliceFlag{
								Name: sectionsFlag,
								Value: cli.NewStringSlice(
									//planners.TitleSection,
									//common.AnnualSection,
									//common.QuarterliesSection,
									//common.MonthliesSection,
									//common.WeekliesSection,
									//common.DailiesSection,
									//common.DailyNotesSection,
									common.DailyReflectSection,
									//common.ToDoSection,
									//common.NotesSection,
								),
							},
						},
						Action: func(appContext *cli.Context) error {
							device, err := devices.New(appContext.String(deviceNameFlag))
							if err != nil {
								return fmt.Errorf("new device: %w", err)
							}

							hand := common.RightHand
							if appContext.String(handFlag) == "left" {
								hand = common.LeftHand
							}

							params := common.NewParams(
								common.ParamWithYear(appContext.Int(yearFlag)),
								common.ParamWithDevice(device),
								common.ParamWithSections(appContext.StringSlice(sectionsFlag)),
								common.ParamWithWeekday(time.Weekday(appContext.Int(weekdayFlag))),
							)

							planner, err := planners.New(planners.MonthsOnSidesTemplate, params)
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
