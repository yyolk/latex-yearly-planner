package main

import (
	"fmt"
	"os"

	"github.com/kudrykv/latex-yearly-planner/lib/planners"
)

func main() {
	planner := planners.New(planners.Params{})
	if err := planner.GenerateFiles("out"); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "generate files: %w", err)
		os.Exit(1)
	}

	if err := planner.Compile("out"); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "compile: %w", err)
		os.Exit(1)
	}
}
