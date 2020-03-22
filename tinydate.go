package tinydate

const tinyDateBinaryVersion byte = 1

const iso8601Date = "2006-01-02"

// TinyDate -
type TinyDate struct {
	year uint16
	// stored as 0-11
	month uint8
	// stored 0-30, depending on the month
	day uint8
}
