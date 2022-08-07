package types

type Layout struct {
	Paper       Paper
	Margin      Margin
	MarginNotes MarginNotes
	Sizes       Sizes

	Debug Debug

	Misc any
}

type Debug struct {
	ShowLinks  bool
	ShowFrames bool
}
