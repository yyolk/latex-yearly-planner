package planners

type MOSParameters struct{}

type mos struct {
	params MOSParameters
}

func (m mos) GenerateFiles(dir string) error {
	panic("implement me")
}

func (m mos) Compile(dir string) error {
	panic("implement me")
}
