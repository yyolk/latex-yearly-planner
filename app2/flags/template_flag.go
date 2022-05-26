package flags

import (
	"errors"
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/templates"
	"github.com/urfave/cli/v2"
)

type TemplateFlag struct {
	cli.StringFlag
}

var UnknownTemplateNameErr = errors.New("unknown template")

func NewTemplateFlag() *TemplateFlag {
	return &TemplateFlag{
		StringFlag: cli.StringFlag{
			Name:     "template-name",
			Required: true,
		},
	}
}

func (r TemplateFlag) Template() (templates.Template, error) {
	switch r.GetValue() {
	default:
		return nil, fmt.Errorf("%s: %w", r.GetValue(), UnknownTemplateNameErr)
	}
}
