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

func TestConstructor(t *testing.T) {
	array := []string{ "dude" }
	aList := New[string, any](array)

	error, isEmpty := aList.IsEmpty()

	if error != nil {
		t.Error("TestConstructor got an error")
		t.Error(error)
	}
	if isEmpty {
		t.Error("list.IsEmpty() in TestConstructor returned false")
	}	
}

func TestListDotMap(t *testing.T) {
	arr := []int { 0, 1, 2, 3, 4 }
	list := New[int, any](arr)

	mapper := func (value int, index int, array *[]int) int {
		return value + 1
	}

	mapped := list.Map(mapper)
	if (*mapped.Array)[4] != 5 {
		t.Error("unexpected value in `mapped[4]`")		
	}
}

func TestListDotGet(t *testing.T) {
	arr := []int { 0, 1, 2, 3, 4 }
	list := New[int, any](arr)

	error, valueAtFour := list.Get(4)
	if error != nil {
		t.Error("error during TestListDotGet")
		t.Error(error)
	}
	if valueAtFour != 4 {
		t.Error("unexpected value in `mapped[4]`")		
	}
}