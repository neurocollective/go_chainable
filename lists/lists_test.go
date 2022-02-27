package lists

import (
	"testing"
)

func TestNewEmptyConstructor(t *testing.T) {
	aList := NewEmpty[string, any]()

	error, isEmpty := aList.IsEmpty()

	if error != nil {
		t.Error("list.IsEmpty() got an error")
		t.Error(error)
	}
	if !isEmpty {
		t.Error("list.IsEmpty() returned false")
	}
}