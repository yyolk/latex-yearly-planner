package ref

type Item struct {
	typ      int
	name     string
	refer    string
	referred bool
}

const (
	text = iota
	reflect
	note
)

func NewText(name, refer string) Item {
	return Item{typ: text, name: name, refer: refer}
}

func NewNote(text, refer string) Item {
	return Item{typ: note, name: text, refer: refer}
}

func NewReflect(text, refer string) Item {
	return Item{typ: reflect, name: text, refer: refer}
}

func (r Item) Build() string {
	if r.referred {
		return NewTargetWithRef(r.name, r.ref()).Build()
	}

	return NewLinkWithRef(r.name, r.ref()).Build()
}

func (r Item) ref() string {
	refer := r.refer
	if refer == "" {
		refer = r.name
	}

	return r.prefix() + refer
}

func (r Item) prefix() string {
	switch r.typ {
	case text:
		return ""
	case note:
		return "note"
	case reflect:
		return "reflect"
	default:
		return "unknown"
	}
}

func (r Item) Ref() Item {
	r.referred = true

	return r
}
