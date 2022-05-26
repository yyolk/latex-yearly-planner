package planners

import "github.com/kudrykv/latex-yearly-planner/app2/devices"

type MonthsOnSides struct {
	params Params
}

func (r *MonthsOnSides) GenerateFor(device devices.Device) error {
	return nil
}
