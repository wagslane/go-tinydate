package tinydate

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestYear(t *testing.T) {
	td := TinyDate{year: 2020}
	assert.Equal(t, 2020, td.Year())
}

func TestMonth(t *testing.T) {
	td := TinyDate{month: 11}
	assert.Equal(t, time.December, td.Month())
}

func TestDay(t *testing.T) {
	td := TinyDate{day: 1}
	assert.Equal(t, 2, td.Day())
}

func TestString(t *testing.T) {
	td, err := New(2017, 11, 4)
	assert.Nil(t, err)
	assert.Equal(t, "2017-11-04", td.String())
}

func TestIsZero(t *testing.T) {
	assert.True(t, TinyDate{}.IsZero())
	assert.True(t, TinyDate{year: 0, month: 0, day: 0}.IsZero())
	assert.False(t, TinyDate{year: 1}.IsZero())
}

func TestToUnix(t *testing.T) {
	td, err := New(2017, 11, 4)
	assert.Nil(t, err)
	assert.Equal(t, int64(1509753600), td.Unix())

	td, err = New(0, 1, 1)
	assert.Nil(t, err)
	assert.Equal(t, int64(-62167219200), td.Unix())

	td, err = New(1970, 1, 1)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), td.Unix())
}

func TestToUnixNano(t *testing.T) {
	td, err := New(2017, 11, 4)
	assert.Nil(t, err)
	assert.Equal(t, int64(1509753600000000000), td.UnixNano())

	td, err = New(1970, 1, 1)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), td.UnixNano())
}

func TestWeekday(t *testing.T) {
	td, err := New(2020, 03, 21)
	assert.Nil(t, err)
	assert.Equal(t, time.Saturday, td.Weekday())
}

func TestYearDay(t *testing.T) {
	td, err := New(2020, 01, 01)
	assert.Nil(t, err)
	assert.Equal(t, 1, td.YearDay())

	td, err = New(2020, 02, 01)
	assert.Nil(t, err)
	assert.Equal(t, 32, td.YearDay())

	td, err = New(2020, 02, 02)
	assert.Nil(t, err)
	assert.Equal(t, 33, td.YearDay())
}

func TestISOWeek(t *testing.T) {
	td, err := New(2020, 03, 21)
	assert.Nil(t, err)
	year, week := td.ISOWeek()
	assert.Equal(t, 2020, year)
	assert.Equal(t, 12, week)
}
func TestDate(t *testing.T) {
	td, err := New(2020, 03, 21)
	assert.Nil(t, err)
	year, month, day := td.Date()
	assert.Equal(t, 2020, year)
	assert.Equal(t, time.March, month)
	assert.Equal(t, 21, day)
}
