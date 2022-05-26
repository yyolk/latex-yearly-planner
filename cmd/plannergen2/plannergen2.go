package main

import (
	"os"

	"github.com/kudrykv/latex-yearly-planner/app2"
)

func main() {
	if err := app2.New(os.Stdin, os.Stdout, os.Stderr).Run(os.Args); err != nil {
		os.Exit(1)
	}
}
