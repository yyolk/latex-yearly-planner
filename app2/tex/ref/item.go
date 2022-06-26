package ref

type Item struct {
	typ      int
	name     string
	refer    string
	referred bool
}

const (
	Text = iota
	Reflect
	Note
	ToDo
)

func NewText(name, refer string) Item {
	return Item{typ: Text, name: name, refer: refer}
}

func NewNote(text, refer string) Item {
	return Item{typ: Note, name: text, refer: refer}
}

func NewReflect(text, refer string) Item {
	return Item{typ: Reflect, name: text, refer: refer}
}

func NewToDo(text, refer string) Item {
	return Item{typ: ToDo, name: text, refer: refer}
}

func NewItem(typ int, name, refer string) Item {
	return Item{typ: typ, name: name, refer: refer}
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
	case Text:
		return ""
	case Note:
		return "note"
	case Reflect:
		return "reflect"
	case ToDo:
		return "todo"
	default:
		return "unknown"
	}
}

func (r Item) Ref() Item {
	r.referred = true

	return r
}
