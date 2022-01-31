package tinydate

import (
	"testing"
)

func TestMarshalBinary(t *testing.T) {
	td := TinyDate{year: 2020, month: 1, day: 3}
	dat, err := td.MarshalBinary()
	assertNil(t, err)
	assertEqual(t, []byte{
		tinyDateBinaryVersion,
		byte(7),
		byte(228),
		byte(1),
		byte(3),
	}, dat)

	newTD := TinyDate{}
	err = newTD.UnmarshalBinary(dat)
	assertNil(t, err)
	assertEqual(t, td, newTD)
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
	assertNil(t, err)
	assertEqual(t, TinyDate{year: 2020, month: 1, day: 3}, td)

	newDat, _ := td.MarshalBinary()
	assertEqual(t, dat, newDat)
}
