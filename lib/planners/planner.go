package planners

type Planner struct{}

type Params struct{}

func New(_ Params) Planner {
	return Planner{}
}

func (p Planner) GenerateFiles(_ string) error {
	panic("not implemented")
}

func (p Planner) Compile(_ string) error {
	panic("not implemented")
}
