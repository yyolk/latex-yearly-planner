package main

import (
	"os"

	"github.com/kudrykv/latex-yearly-planner/app2"
)

func main() {
	if err := app2.New().Run(os.Args); err != nil {
		os.Exit(1)
	}
}
