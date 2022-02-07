package tinydate

import (
	"testing"
)

func TestMarshalText(t *testing.T) {
	td := TinyDate{year: 2020, month: 1, day: 3}
	dat, err := td.MarshalText()
	assertNil(t, err)
	assertEqual(t, []byte(`2020-02-04`), dat)
}

func TestUnmarshalText(t *testing.T) {
	dat := []byte(`2020-02-04`)
	td := TinyDate{}
	err := td.UnmarshalText(dat)
	assertNil(t, err)
	assertEqual(t, TinyDate{year: 2020, month: 1, day: 3}, td)
}
