package tinydate

import (
	"fmt"
	"math"
	"time"
)

// FromTime converts a time.Time into a TinyDate
// all tinyTinyDate.TinyDates are UTC timezone, so that conversion will
// be made here
func FromTime(t time.Time) (TinyDate, error) {
	t = t.UTC()
	year := t.Year()
	if year > math.MaxUint16 {
		return TinyDate{}, fmt.Errorf("tinydate FromTime: year must be less than %v", math.MaxUint16)
	}
	return TinyDate{
		year:  uint16(year),
		month: uint8(t.Month() - 1),
		day:   uint8(t.Day() - 1),
	}, nil
}

// ToTime converts a tinyTinyDate.TinyDate to a time.Time, always UTC
func (td TinyDate) ToTime() time.Time {
	return time.Date(int(td.year), time.Month(td.Month()), int(td.Day()), 0, 0, 0, 0, time.UTC)
}
