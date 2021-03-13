package tinydate

import (
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestSizes(t *testing.T) {
	assert.Equal(t, uintptr(4), unsafe.Sizeof(TinyDate{}))

	tm := time.Now().UTC()
	assert.Equal(t, uintptr(24), unsafe.Sizeof(tm))
}
