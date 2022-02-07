package tinydate

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, object1, object2 interface{}) {
	if !reflect.DeepEqual(object1, object2) {
		t.Helper()
		t.Errorf("expected %v and %v to be equal, but they are not equal", object1, object2)
	}
}

func assertNil(t *testing.T, object interface{}) {
	if !isNil(object) {
		t.Helper()
		t.Errorf("expected %v to be nil, but wasn't nil", object)
	}
}

func assertNotNil(t *testing.T, object interface{}) {
	if isNil(object) {
		t.Helper()
		t.Errorf("expected %v not to be nil, but was nil", object)
	}
}

func assertTrue(t *testing.T, b bool) {
	t.Helper()
	assertEqual(t, b, true)
}

func assertFalse(t *testing.T, b bool) {
	t.Helper()
	assertEqual(t, b, false)
}

func isNil(object interface{}) bool {
	if object == nil {
		return true
	}
	return reflect.ValueOf(object).IsNil()
}
