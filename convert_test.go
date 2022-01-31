package tinydate

import (
	"testing"
	"time"
)

func TestFromTimeValid(t *testing.T) {
	tm := time.Date(2020, time.Month(5), 2, 0, 0, 0, 0, time.UTC)
	td, err := FromTime(tm)
	assertNil(t, err)
	assertEqual(t, TinyDate{year: 2020, month: 4, day: 1}, td)
}

func TestFromTimeInvalid(t *testing.T) {
	tm := time.Date(100000, time.Month(5), 2, 0, 0, 0, 0, time.UTC)
	_, err := FromTime(tm)
	assertNotNil(t, err)
}

func TestToTime(t *testing.T) {
	td := TinyDate{year: 2020, month: 4, day: 2}
	tm := td.ToTime()
	assertEqual(t, time.Date(2020, 5, 3, 0, 0, 0, 0, time.UTC), tm)

	td = TinyDate{year: 1970, month: 3, day: 1}
	tm = td.ToTime()
	assertEqual(t, time.Date(1970, 4, 2, 0, 0, 0, 0, time.UTC), tm)
}
