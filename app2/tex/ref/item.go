package ref

type Item struct {
	typ      int
	text     string
	refer    string
	referred bool
}

const (
	Text = iota
	Note
)

func NewText(text, refer string) Item {
	return Item{typ: Text, text: text, refer: refer}
}

func NewNote(text, refer string) Item {
	return Item{typ: Note, text: text, refer: refer}
}

func (r Item) Build() string {
	if r.referred {
		return NewTargetWithRef(r.text, r.ref()).Build()
	}

	return NewLinkWithRef(r.text, r.ref()).Build()
}

func (r Item) ref() string {
	refer := r.refer
	if refer == "" {
		refer = r.text
	}

	return r.prefix() + refer
}

func (r Item) prefix() string {
	switch r.typ {
	case Text:
		return ""
	case Note:
		return "note"
	default:
		return "unknown"
	}
}

func (r Item) Ref() Item {
	r.referred = true

	return r
}
