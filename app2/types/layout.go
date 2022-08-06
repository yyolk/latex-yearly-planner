package types

type Layout[T any] struct {
	Paper       Paper
	Margin      Margin
	MarginNotes MarginNotes
	Sizes       Sizes

	Misc T
}
