package sections

import (
	"errors"
	"fmt"
)

type IndexParameters struct {
	ItemsPerPage int
	Pages        int
}

var (
	ErrNonPositiveItemsPerPage = errors.New("negative index per page")
	ErrNonPositivePages        = errors.New("negative pages")
)

func (r IndexParameters) Test() error {
	if r.ItemsPerPage <= 0 {
		return ErrNonPositiveItemsPerPage
	}

	if r.Pages <= 0 {
		return ErrNonPositivePages
	}

	return nil
}

type Index struct {
	parameters IndexParameters
}

func NewIndex(parameters IndexParameters) (Index, error) {
	if err := parameters.Test(); err != nil {
		return Index{}, fmt.Errorf("test index parameters: %w", err)
	}

	return Index{
		parameters: parameters,
	}, nil
}

func (r Index) Build() ([]string, error) {
	pages := make([]string, 0, r.parameters.Pages)

	for i := 0; i < r.parameters.Pages; i++ {
		pages = append(pages, fmt.Sprintf("index %d", i))
	}

	return pages, nil
}
