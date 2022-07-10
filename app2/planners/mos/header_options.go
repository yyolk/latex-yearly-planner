package mos

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/app2/tex/cell"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type headerOption func(*header)

func headerWithTexYear(year texcalendar.Year) headerOption {
	return func(header *header) {
		header.year = year
	}
}

func headerWithHand(hand common.MainHand) headerOption {
	return func(header *header) {
		header.hand = hand
	}
}

func headerWithTitle(left string) headerOption {
	return func(header *header) {
		header.title = left
	}
}

func headerWithActions(cells cell.Cells) headerOption {
	return func(header *header) {
		header.action = cells
	}
}

func headerAddAction(cell cell.Cell) headerOption {
	return func(header *header) {
		if header.hand == common.LeftHand {
			header.action = header.action.Push(cell)

			return
		}

		header.action = header.action.Shift(cell)
	}
}

func headerSelectQuarter(quarter texcalendar.Quarter) headerOption {
	return func(header *header) {
		header.selectedQuarter = quarter
	}
}

func headerSelectMonths(months ...time.Month) headerOption {
	return func(header *header) {
		header.selectedMonths = months
	}
}
