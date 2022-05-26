package planners

type Params struct {
	Name string
}

func NewParams(name string) Params {
	return Params{Name: name}
}
