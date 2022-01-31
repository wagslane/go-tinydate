package tinydate

import (
	"testing"
	"time"
)

func TestYear(t *testing.T) {
	td := TinyDate{year: 2020}
	assertEqual(t, 2020, td.Year())
}

func TestMonth(t *testing.T) {
	td := TinyDate{month: 11}
	assertEqual(t, time.December, td.Month())
}

func TestDay(t *testing.T) {
	td := TinyDate{day: 1}
	assertEqual(t, 2, td.Day())
}

func TestString(t *testing.T) {
	td, err := New(2017, 11, 4)
	assertNil(t, err)
	assertEqual(t, "2017-11-04", td.String())
}

func TestIsZero(t *testing.T) {
	assertTrue(t, TinyDate{}.IsZero())
	assertTrue(t, TinyDate{year: 0, month: 0, day: 0}.IsZero())
	assertFalse(t, TinyDate{year: 1}.IsZero())
}

func TestToUnix(t *testing.T) {
	td, err := New(2017, 11, 4)
	assertNil(t, err)
	assertEqual(t, int64(1509753600), td.Unix())

	td, err = New(0, 1, 1)
	assertNil(t, err)
	assertEqual(t, int64(-62167219200), td.Unix())

	td, err = New(1970, 1, 1)
	assertNil(t, err)
	assertEqual(t, int64(0), td.Unix())
}

func TestToUnixNano(t *testing.T) {
	td, err := New(2017, 11, 4)
	assertNil(t, err)
	assertEqual(t, int64(1509753600000000000), td.UnixNano())

	td, err = New(1970, 1, 1)
	assertNil(t, err)
	assertEqual(t, int64(0), td.UnixNano())
}

func TestWeekday(t *testing.T) {
	td, err := New(2020, 03, 21)
	assertNil(t, err)
	assertEqual(t, time.Saturday, td.Weekday())
}

func TestYearDay(t *testing.T) {
	td, err := New(2020, 01, 01)
	assertNil(t, err)
	assertEqual(t, 1, td.YearDay())

	td, err = New(2020, 02, 01)
	assertNil(t, err)
	assertEqual(t, 32, td.YearDay())

	td, err = New(2020, 02, 02)
	assertNil(t, err)
	assertEqual(t, 33, td.YearDay())
}

func TestISOWeek(t *testing.T) {
	td, err := New(2020, 03, 21)
	assertNil(t, err)
	year, week := td.ISOWeek()
	assertEqual(t, 2020, year)
	assertEqual(t, 12, week)
}
func TestDate(t *testing.T) {
	td, err := New(2020, 03, 21)
	assertNil(t, err)
	year, month, day := td.Date()
	assertEqual(t, 2020, year)
	assertEqual(t, time.March, month)
	assertEqual(t, 21, day)
}
