package app

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/kudrykv/latex-yearly-planner/app/planners"
	"github.com/kudrykv/latex-yearly-planner/app/planners/mos"
	"github.com/urfave/cli/v2"
)

type App struct {
	app *cli.App
}

const (
	parametersPathFlag = "parameters-path"
)

func New(reader io.Reader, writer, errWriter io.Writer) *App {
	return (&App{}).setupCli(reader, writer, errWriter)
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
							&cli.StringFlag{Name: parametersPathFlag},
						},

						Action: func(appContext *cli.Context) error {
							var parameters mos.Parameters
							if err := readToml(appContext.String(parametersPathFlag), &parameters); err != nil {
								return fmt.Errorf("read parameters: %w", err)
							}

							planner, err := planners.New("mos", parameters)
							if err != nil {
								return fmt.Errorf("new planner: %w", err)
							}

							if err = planner.Generate(); err != nil {
								return fmt.Errorf("generate: %w", err)
							}

							if err = planner.WriteTeXTo("./out"); err != nil {
								return fmt.Errorf("write tex: %w", err)
							}

							if err = planner.Compile(appContext.Context, "./out"); err != nil {
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

func readToml(path string, dst any) error {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	if err = toml.Unmarshal(fileBytes, dst); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	return nil
}

func (r *App) Run(ctx context.Context, args []string) error {
	return r.app.RunContext(ctx, args)
}
