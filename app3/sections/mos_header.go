package sections

type MOSHeader struct {
}

func NewMOSHeader() MOSHeader {
	return MOSHeader{}
}

func (M MOSHeader) Build() ([]string, error) {
	return []string{"header" + "\n\n"}, nil
}
