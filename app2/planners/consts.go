package planners

import (
	"errors"
)

const (
	TitleSection       = "title"
	AnnualSection      = "annual"
	QuarterliesSection = "quarterlies"
	MonthliesSection   = "monthlies"
	WeekliesSection    = "weeklies"
	DailiesSection     = "dailies"
)

var UnknownDeviceTypeErr = errors.New("unknown device type")
