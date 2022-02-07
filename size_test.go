package tinydate

import (
	"testing"
	"time"
	"unsafe"
)

func TestSizes(t *testing.T) {
	assertEqual(t, uintptr(4), unsafe.Sizeof(TinyDate{}))

	tm := time.Now().UTC()
	assertEqual(t, uintptr(24), unsafe.Sizeof(tm))
}
