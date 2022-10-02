package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kudrykv/latex-yearly-planner/app"
)

func main() {
	if err := app.New(os.Stdin, os.Stdout, os.Stderr).Run(context.Background(), os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
