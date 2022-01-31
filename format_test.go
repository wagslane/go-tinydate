package tinydate

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	td, err := Parse(time.RFC3339, "2020-03-05T00:00:00Z")
	assertNil(t, err)
	assertEqual(t, TinyDate{year: 2020, month: 2, day: 4}, td)

	td, err = Parse(iso8601Date, "2020-03-05")
	assertNil(t, err)
	assertEqual(t, TinyDate{year: 2020, month: 2, day: 4}, td)
}

func TestParseInLocation(t *testing.T) {
	td, err := ParseInLocation(time.RFC3339, "2020-03-05T00:00:00Z", time.UTC)
	assertNil(t, err)
	assertEqual(t, TinyDate{year: 2020, month: 2, day: 4}, td)

	td, err = ParseInLocation(iso8601Date, "2020-03-05", time.UTC)
	assertNil(t, err)
	assertEqual(t, TinyDate{year: 2020, month: 2, day: 4}, td)
}

func TestAppendFormat(t *testing.T) {
	td, err := New(2017, 11, 4)
	assertNil(t, err)
	text := []byte("Date: ")

	text = td.AppendFormat(text, iso8601Date)
	assertEqual(t, "Date: 2017-11-04", string(text))
}

func TestFormat(t *testing.T) {
	td, err := New(2017, 11, 4)
	assertNil(t, err)

	assertEqual(t, "2017-11-04", string(td.Format(iso8601Date)))
	assertEqual(t, "2017-11-04T00:00:00Z", string(td.Format(time.RFC3339)))
}
