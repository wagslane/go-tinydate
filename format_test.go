package tinydate

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	td, err := Parse(time.RFC3339, "2020-03-05T00:00:00Z")
	assert.Nil(t, err)
	assert.Equal(t, TinyDate{year: 2020, month: 2, day: 4}, td)

	td, err = Parse(iso8601Date, "2020-03-05")
	assert.Nil(t, err)
	assert.Equal(t, TinyDate{year: 2020, month: 2, day: 4}, td)
}

func TestParseInLocation(t *testing.T) {
	td, err := ParseInLocation(time.RFC3339, "2020-03-05T00:00:00Z", time.UTC)
	assert.Nil(t, err)
	assert.Equal(t, TinyDate{year: 2020, month: 2, day: 4}, td)

	td, err = ParseInLocation(iso8601Date, "2020-03-05", time.UTC)
	assert.Nil(t, err)
	assert.Equal(t, TinyDate{year: 2020, month: 2, day: 4}, td)
}

func TestAppendFormat(t *testing.T) {
	td, err := New(2017, 11, 4)
	assert.Nil(t, err)
	text := []byte("Date: ")

	text = td.AppendFormat(text, iso8601Date)
	assert.Equal(t, "Date: 2017-11-04", string(text))
}

func TestFormat(t *testing.T) {
	td, err := New(2017, 11, 4)
	assert.Nil(t, err)

	assert.Equal(t, "2017-11-04", string(td.Format(iso8601Date)))
	assert.Equal(t, "2017-11-04T00:00:00Z", string(td.Format(time.RFC3339)))
}
