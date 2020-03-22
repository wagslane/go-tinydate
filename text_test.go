package tinydate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalText(t *testing.T) {
	td := TinyDate{year: 2020, month: 1, day: 3}
	dat, err := td.MarshalText()
	assert.Nil(t, err)
	assert.Equal(t, []byte(`2020-02-04`), dat)
}

func TestUnmarshalText(t *testing.T) {
	dat := []byte(`2020-02-04`)
	td := TinyDate{}
	err := td.UnmarshalText(dat)
	assert.Nil(t, err)
	assert.Equal(t, TinyDate{year: 2020, month: 1, day: 3}, td)
}
