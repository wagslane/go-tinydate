package tinydate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalBinary(t *testing.T) {
	td := TinyDate{year: 2020, month: 1, day: 3}
	dat, err := td.MarshalBinary()
	assert.Nil(t, err)
	assert.Equal(t, []byte{
		tinyDateBinaryVersion,
		byte(7),
		byte(228),
		byte(1),
		byte(3),
	}, dat)

	newTD := TinyDate{}
	err = newTD.UnmarshalBinary(dat)
	assert.Nil(t, err)
	assert.Equal(t, td, newTD)
}

func TestUnmarshalBinary(t *testing.T) {
	dat := []byte{
		tinyDateBinaryVersion,
		byte(7),
		byte(228),
		byte(1),
		byte(3),
	}
	td := TinyDate{}
	err := td.UnmarshalBinary(dat)
	assert.Nil(t, err)
	assert.Equal(t, TinyDate{year: 2020, month: 1, day: 3}, td)

	newDat, _ := td.MarshalBinary()
	assert.Equal(t, dat, newDat)
}
