package test

import (
	"fmt"
	"os"
	"path"
)

func Fixture(filename string) string {
	pathToFile := path.Join("fixtures", fmt.Sprintf("%s.tex", filename))

	fileBytes, err := os.ReadFile(pathToFile)
	if err != nil {
		panic(fmt.Errorf("read fixture %s: %w", pathToFile, err))
	}

	return string(fileBytes)
}
