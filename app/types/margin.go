package types

type Margin struct {
	Top    Millimeters
	Right  Millimeters
	Bottom Millimeters
	Left   Millimeters
}

type MarginNotes struct {
	Width     Millimeters
	Separator Millimeters
	Reverse   bool
}
