package tinydate

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAfter(t *testing.T) {
	td := TinyDate{year: 2020}
	tu := TinyDate{year: 2020}
	assert.False(t, td.After(tu))

	td = TinyDate{year: 2020, month: 5, day: 10}
	tu = TinyDate{year: 2020, month: 5, day: 10}
	assert.False(t, td.After(tu))

	td = TinyDate{year: 2021, month: 5, day: 10}
	tu = TinyDate{year: 2020, month: 5, day: 10}
	assert.True(t, td.After(tu))

	td = TinyDate{year: 2020, month: 6, day: 10}
	tu = TinyDate{year: 2020, month: 5, day: 10}
	assert.True(t, td.After(tu))

	td = TinyDate{year: 2020, month: 5, day: 11}
	tu = TinyDate{year: 2020, month: 5, day: 10}
	assert.True(t, td.After(tu))

	td = TinyDate{year: 2020, month: 4, day: 11}
	tu = TinyDate{year: 2020, month: 5, day: 10}
	assert.False(t, td.After(tu))

	td = TinyDate{year: 2019, month: 6, day: 10}
	tu = TinyDate{year: 2020, month: 5, day: 10}
	assert.False(t, td.After(tu))
}

func TestBefore(t *testing.T) {
	td := TinyDate{year: 2020}
	tu := TinyDate{year: 2020}
	assert.False(t, td.Before(tu))

	td = TinyDate{year: 2020, month: 5, day: 10}
	tu = TinyDate{year: 2020, month: 5, day: 10}
	assert.False(t, td.Before(tu))

	td = TinyDate{year: 2021, month: 5, day: 10}
	tu = TinyDate{year: 2020, month: 5, day: 10}
	assert.False(t, td.Before(tu))

	td = TinyDate{year: 2020, month: 6, day: 10}
	tu = TinyDate{year: 2020, month: 5, day: 10}
	assert.False(t, td.Before(tu))

	td = TinyDate{year: 2020, month: 5, day: 11}
	tu = TinyDate{year: 2020, month: 5, day: 10}
	assert.False(t, td.Before(tu))

	td = TinyDate{year: 2020, month: 4, day: 11}
	tu = TinyDate{year: 2020, month: 5, day: 10}
	assert.True(t, td.Before(tu))

	td = TinyDate{year: 2019, month: 6, day: 10}
	tu = TinyDate{year: 2020, month: 5, day: 10}
	assert.True(t, td.Before(tu))
}

func TestAdd(t *testing.T) {
	td := TinyDate{year: 2020, month: 5, day: 10}
	assert.Equal(t,
		TinyDate{year: 2020, month: 5, day: 11},
		td.Add(time.Hour*24),
	)

	td = TinyDate{year: 2020, month: 5, day: 10}
	assert.Equal(t,
		TinyDate{year: 2020, month: 5, day: 9},
		td.Add(-time.Hour*24),
	)

	td = TinyDate{year: 2020, month: 12, day: 31}
	assert.Equal(t,
		TinyDate{year: 2021, month: 1, day: 1},
		td.Add(time.Hour*24),
	)
}

func TestAddDate(t *testing.T) {
	td := TinyDate{year: 2020, month: 5, day: 10}
	assert.Equal(t,
		TinyDate{year: 2020, month: 5, day: 11},
		td.AddDate(0, 0, 1),
	)

	td = TinyDate{year: 2020, month: 5, day: 10}
	assert.Equal(t,
		TinyDate{year: 2020, month: 5, day: 9},
		td.AddDate(0, 0, -1),
	)

	td = TinyDate{year: 2020, month: 12, day: 31}
	assert.Equal(t,
		TinyDate{year: 2021, month: 1, day: 1},
		td.AddDate(0, 0, 1),
	)

	td = TinyDate{year: 2011, month: 01, day: 01}
	assert.Equal(t,
		TinyDate{year: 2010, month: 03, day: 04},
		td.AddDate(-1, 2, 3),
	)

	td = TinyDate{year: 2020, month: 01, day: 01}
	assert.Equal(t,
		TinyDate{year: 2021, month: 01, day: 01},
		td.AddDate(1, 0, 0),
	)
}

func TestEqual(t *testing.T) {
	td, err := New(2020, 03, 04)
	assert.Nil(t, err)
	assert.True(t, td.Equal(td))

	td, err = New(2020, 03, 04)
	assert.Nil(t, err)
	tu, err := New(2020, 03, 05)
	assert.Nil(t, err)
	assert.False(t, td.Equal(tu))

	td, err = New(2020, 03, 04)
	assert.Nil(t, err)
	tu, err = New(2020, 04, 04)
	assert.Nil(t, err)
	assert.False(t, td.Equal(tu))

	td, err = New(2021, 03, 04)
	assert.Nil(t, err)
	tu, err = New(2020, 03, 04)
	assert.Nil(t, err)
	assert.False(t, td.Equal(tu))
}

func TestSub(t *testing.T) {
	td, err := New(2020, 03, 05)
	assert.Nil(t, err)
	tu, err := New(2020, 03, 04)
	assert.Nil(t, err)
	assert.Equal(t, time.Hour*24, td.Sub(tu))

	td, err = New(2020, 03, 05)
	assert.Nil(t, err)
	tu, err = New(2020, 03, 05)
	assert.Nil(t, err)
	assert.Equal(t, time.Duration(0), td.Sub(tu))

	td, err = New(2020, 03, 05)
	assert.Nil(t, err)
	tu, err = New(2020, 03, 06)
	assert.Nil(t, err)
	assert.Equal(t, -time.Hour*24, td.Sub(tu))
}
