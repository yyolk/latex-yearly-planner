package mos

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/app3/sections"
	"github.com/kudrykv/latex-yearly-planner/app3/types"
)

type Parameters struct {
	Sections []string

	Document types.Document

	Year                   int
	Weekday                time.Weekday
	MOSHeaderParameters    sections.MOSHeaderParameters
	DailyParameters        sections.DailyParameters
	DailyNotesParameters   sections.DailyNotesParameters
	DailyReflectParameters sections.DailyReflectParameters
	NotesParameters        sections.NotesParameters
}

func (r Parameters) DailyNotesEnabled() bool {
	for _, section := range r.Sections {
		if section == "daily_notes" {
			return true
		}
	}

	return false
}

func (r Parameters) ReflectEnabled() bool {
	for _, section := range r.Sections {
		if section == "daily_reflect" {
			return true
		}
	}

	return false
}
