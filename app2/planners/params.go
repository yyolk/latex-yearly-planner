package planners

type Params struct {
	Name         string
	TemplateData TemplateData
}

func NewParams(name string) Params {
	return Params{
		Name:         name,
		TemplateData: TemplateData{},
	}
}
