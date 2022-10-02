package mos

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/app/sections"
	"github.com/kudrykv/latex-yearly-planner/app/types"
)

type Parameters struct {
	Sections []string

	Document types.Document

	Year                   int
	Weekday                time.Weekday
	MOSHeaderParameters    sections.MOSHeaderParameters
	AnnualParameters       sections.AnnualParameters
	QuarterlyParameters    sections.QuarterlyParameters
	MonthlyParameters      sections.MonthlyParameters
	WeeklyParameters       sections.WeeklyParameters
	DailyParameters        sections.DailyParameters
	DailyNotesParameters   sections.DailyNotesParameters
	DailyReflectParameters sections.DailyReflectParameters
	NotesParameters        sections.NotesParameters
	TodosParameters        sections.TodosParameters
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

func (r Parameters) NotesEnabled() bool {
	for _, section := range r.Sections {
		if section == "notes" {
			return true
		}
	}

	return false
}

func (r Parameters) TodosEnabled() bool {
	for _, section := range r.Sections {
		if section == "todo" {
			return true
		}
	}

	return false
}
