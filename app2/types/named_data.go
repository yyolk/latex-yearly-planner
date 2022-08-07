package types

type NamedData struct {
	Name string
	Data []byte
}

type NamedDatas []NamedData

func (r *NamedDatas) Append(name string, data []byte) {
	*r = append(*r, NamedData{name, data})
}
