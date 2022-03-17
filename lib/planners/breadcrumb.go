package planners

type BreadcrumbParams struct{}

type breadcrumb struct {
	Params BreadcrumbParams
}

func (r breadcrumb) GenerateFiles(_ string) error {
	panic("not implemented")
}

func (r breadcrumb) Compile(_ string) error {
	panic("not implemented")
}
