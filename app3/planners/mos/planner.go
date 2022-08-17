package mos

import (
	"github.com/kudrykv/latex-yearly-planner/app3/types"
)

type Planner struct {
}

func New(Parameters) *Planner {
	return &Planner{}
}

func (r *Planner) Generate() (types.NamedBuffers, error) {
	panic("not implemented")
}
