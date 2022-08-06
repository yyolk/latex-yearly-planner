package types

type Margin struct {
	Top    Millimeters
	Right  Millimeters
	Bottom Millimeters
	Left   Millimeters
}

type MarginNotes struct {
	Margin  Millimeters
	Width   Millimeters
	Reverse bool
}
