package tinydate

import (
	"encoding/json"
	"testing"
)

func TestMarshalJSON(t *testing.T) {
	dat, err := TinyDate{year: 2020, month: 1, day: 3}.MarshalJSON()
	assertNil(t, err)
	assertEqual(t, `"2020-02-04"`, string(dat))

	s := struct {
		TD TinyDate
	}{
		TD: TinyDate{year: 2020, month: 1, day: 3},
	}
	dat, err = json.Marshal(s)
	assertNil(t, err)
	assertEqual(t, `{"TD":"2020-02-04"}`, string(dat))
}

func TestUnmarshalJSON(t *testing.T) {
	dat := []byte(`"2020-02-04"`)
	td := TinyDate{}
	err := td.UnmarshalJSON(dat)
	assertNil(t, err)
	assertEqual(t, TinyDate{year: 2020, month: 1, day: 3}, td)

	dat = []byte(`{"TD":"2020-02-04"}`)
	s := struct {
		TD TinyDate
	}{}
	err = json.Unmarshal(dat, &s)
	assertNil(t, err)
	assertEqual(t, struct {
		TD TinyDate
	}{
		TD: TinyDate{year: 2020, month: 1, day: 3},
	}, s)
}
