package planners

type BreadcrumbParams struct{}

type Breadcrumb struct {
	Params BreadcrumbParams
}

func (r Breadcrumb) GenerateFiles(_ string) error {
	panic("not implemented")
}

func (r Breadcrumb) Compile(_ string) error {
	panic("not implemented")
}
