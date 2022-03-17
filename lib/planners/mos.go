package planners

type MOSParams struct{}

type mos struct {
	params MOSParams
}

func (m mos) GenerateFiles(dir string) error {
	panic("implement me")
}

func (m mos) Compile(dir string) error {
	panic("implement me")
}
