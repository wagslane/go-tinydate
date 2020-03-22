package tinydate

import (
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestSizes(t *testing.T) {
	td := TinyDate{}
	assert.Equal(t, uintptr(4), unsafe.Sizeof(td))

	tm := time.Now().UTC()
	assert.Equal(t, uintptr(24), unsafe.Sizeof(tm))
}
