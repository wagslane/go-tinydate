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
