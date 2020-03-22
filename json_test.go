package tinydate

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalJSON(t *testing.T) {
	dat, err := TinyDate{year: 2020, month: 1, day: 3}.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, `"2020-02-04"`, string(dat))

	s := struct {
		TD TinyDate
	}{
		TD: TinyDate{year: 2020, month: 1, day: 3},
	}
	dat, err = json.Marshal(s)
	assert.Nil(t, err)
	assert.Equal(t, `{"TD":"2020-02-04"}`, string(dat))
}

func TestUnmarshalJSON(t *testing.T) {
	dat := []byte(`"2020-02-04"`)
	td := TinyDate{}
	err := td.UnmarshalJSON(dat)
	assert.Nil(t, err)
	assert.Equal(t, TinyDate{year: 2020, month: 1, day: 3}, td)

	dat = []byte(`{"TD":"2020-02-04"}`)
	s := struct {
		TD TinyDate
	}{}
	err = json.Unmarshal(dat, &s)
	assert.Nil(t, err)
	assert.Equal(t, struct {
		TD TinyDate
	}{
		TD: TinyDate{year: 2020, month: 1, day: 3},
	}, s)
}
