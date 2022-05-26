package planners

type Params struct {
	Name         string
	TemplateData TemplateData
}

type TemplateData struct {
	Year int
}

func NewParams(name string) Params {
	return Params{
		Name:         name,
		TemplateData: TemplateData{},
	}
}
