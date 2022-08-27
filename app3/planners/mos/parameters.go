package mos

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/app3/sections"
	"github.com/kudrykv/latex-yearly-planner/app3/types"
)

type Parameters struct {
	Sections []string

	Document types.Document

	Year            int
	Weekday         time.Weekday
	DailyParameters sections.DailyParameters
}
