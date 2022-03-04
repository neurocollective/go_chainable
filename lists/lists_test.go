package lists

import (
	"testing"
)

func TestRaw(t *testing.T) {
	aList := NewEmpty[string, any]()

	array := aList.Raw()

	if len(array) != 0 {
		t.Error("list.Raw() returned unexpected result in TestRaw")
	}
}

func TestRawPointer(t *testing.T) {
	list := NewEmpty[string, any]()

	arrayPtr := list.RawPointer()

	if arrayPtr != list.Array {
		t.Error("list.RawPointer() returned unexpected result in TestRawPointer")
	}
}

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

	mapper := func (value int, index int) int {
		return value + 1
	}

	mapped := list.Map(mapper)
	if (*mapped.Array)[4] != 5 {
		t.Error("unexpected value in `mapped[4]` in TestListDotMap")		
	}
}

func TestListDotMapFull(t *testing.T) {
	arr := []int { 0, 1, 2, 3, 4 }
	list := New[int, any](arr)

	mapper := func (value int, index int, array *[]int) int {
		return value + 1
	}

	mapped := list.MapFull(mapper)
	if (*mapped.Array)[4] != 5 {
		t.Error("unexpected value in `mapped[4]` in TestListDotMapFull")		
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
		t.Error("unexpected value in `mapped[4]` in TestListDotGet")		
	}
}

func TestReduce(t *testing.T) {
	array := []int { 1, 2 }
	list := New[int, int](array)
	added := list.Reduce(func(accumulator int, val int, index int) int {
		return accumulator + val
	}, 0)
	if added != 3 {
		t.Error(".Reduce() in TestReduce returned unexpected value")			
	}
}

func TestReduceFull(t *testing.T) {
	array := []int { 1, 2 }
	list := New[int, int](array)
	added := list.ReduceFull(func(accumulator int, val int, index int, array *[]int) int {
		return accumulator + val
	}, 0)
	if added != 3 {
		t.Error(".Reduce() in TestReduceFull returned unexpected value")			
	}
}

func TestFilter(t *testing.T) {
	array := []int { 1, 2 }
	list := New[int, any](array)
	newList := list.Filter(func(val int, index int) bool {
		return val < 2
	})
	error, value := newList.Get(0)
	if error != nil {
		t.Error(".Filter() in TestFilter returned an error")			
	}
	if value != 1 {
		t.Error(".Filter() in TestFilter returned unexpected value")			
	}	
}

func TestFilterFull(t *testing.T) {
	array := []int { 1, 2 }
	list := New[int, any](array)
	newList := list.FilterFull(func(val int, index int, array *[]int) bool {
		return val < 2
	})
	error, value := newList.Get(0)
	if error != nil {
		t.Error(".Filter() in TestFilterFull returned an error")			
	}
	if value != 1 {
		t.Error(".Filter() in TestFilterFull returned unexpected value")			
	}	
}

func TestChain(t *testing.T) {
	array := []int { 1, 2 }
	list := New[int, int](array)
	added := list.Map(func(val int, index int) int {
		return val + 1
	}).Reduce(func(accumulator int, val int, index int) int {
		return accumulator + val
	}, 0)
	if added != 5 {
		t.Error(".Map().Reduce() in TestChain returned unexpected value")			
	}
}

func TestReduceWithDifferingResultType(t *testing.T) {
	type Number struct {
		Value int
	}

	array := []Number { Number{1}, Number{2} }
	list := New[Number, int](array)
	added := list.Reduce(func(accumulator int, val Number, index int) int {
		return accumulator + val.Value
	}, 0)
	if added != 3 {
		t.Error(".Reduce() in TestReduceWithDifferingResultType returned unexpected value")			
	}	
}

func TestAppend(t *testing.T) {
	list := NewEmpty[int, any]()

	list.Append(&[]int { 1, 2 })
	errorZero, indexZero := list.Get(0)
	errorOne, indexOne := list.Get(1)

	if errorZero != nil || errorOne != nil {
		t.Error("TestAppend returned an error")
		t.Error(errorZero.Error())
		t.Error(errorOne.Error())
	}
	if indexZero != 1 {
		t.Error("indexZero in TestAppend returned unexpected value")		
	}
	if indexOne != 2 {
		t.Error("indexOne in TestAppend returned unexpected value")				
	}
}

func TestAdd(t *testing.T) {
	list := NewEmpty[int, any]()

	list.Add(1)
	errorZero, indexZero := list.Get(0)

	if errorZero != nil {
		t.Error("TestAdd returned an error")
		t.Error(errorZero.Error())
	}
	if indexZero != 1 {
		t.Error("indexZero in TestAdd returned unexpected value")	
	}
}

func TestSetCap(t *testing.T) {
	list := NewEmpty[int, any]()

	list.SetCap(100)

	error, capacity := list.Cap()

	if error != nil {
		t.Error("TestSetCap returned an error")
		t.Error(error.Error())
	}
	if capacity != 100 {
		t.Error("indexZero in TestSetCap returned unexpected value")	
	}
}

func TestIncrementCap(t *testing.T) {
	list := NewEmpty[int, any]()

	list.SetCap(100)

	list.IncrementCap(100)

	error, capacity := list.Cap()

	if error != nil {
		t.Error("TestIncrementCap returned an error")
		t.Error(error.Error())
	}
	if capacity != 200 {
		t.Error("indexZero in TestIncrementCap returned unexpected value")	
	}
}

func TestFind(t *testing.T) {
	list := New[int, any]([]int{ 8, 6, 7, 5, 3, 0, 9 })

	error, target := list.Find(func(element int, index int) bool {
		if element == 0 {
			return true
		}
		return false
	})
	if error != nil {
		t.Error("TestFind returned an error")
	}
	if *target != 0 {
		t.Error("TestFind did not return the expected value")
	}
}

func TestString(t *testing.T) {
	list := New[int, any]([]int{ 1, 2 })

	str := list.String()

	if str != "[1 2]" {
		t.Error("TestString did not return the expected value")
	}
}

func TestIndexOf(t *testing.T) {
	list := New[int, any]([]int{ 1, 2 })

	error, theIndex := list.IndexOf(func(element int) bool {
		if element == 2 {
			return true
		}
		return false
	})
	if error != nil {
		t.Error("TestIndexOf returned an error")		
	}
	if theIndex != 1 {
		t.Error("TestIndexOf did not return the expected value")
	}	
}

func TestSize(t *testing.T) {
	list := New[int, any]([]int{ 1, 2 })

	error, size := list.Size()
	if error != nil {
		t.Error("TestSize returned an error")		
	}
	if size != 2 {
		t.Error("TestSize did not return the expected value")
	}	
}

func TestCap(t *testing.T) {
	list := New[int, any]([]int{ 1, 2 })

	error, cap := list.Cap()
	if error != nil {
		t.Error("TestCap returned an error")
	}
	if cap != 2 {
		t.Error("TestCap did not return the expected value")
	}	
}

func TestIsEmpty(t *testing.T) {
	list := NewEmpty[int, any]()

	error, empty := list.IsEmpty()
	if error != nil {
		t.Error("TestIsEmpty returned an error")		
	}
	if !empty {
		t.Error("TestIsEmpty did not return the expected empty value")
	}	
}

func TestGet(t *testing.T) {
	list := New[int, any]([]int{ 1, 2 })

	error, value := list.Get(1)
	if error != nil {
		t.Error("TestGet returned an error")		
	}
	if value != 2 {
		t.Error("TestGet did not return the expected value")
	}	
}

func TestLast(t *testing.T) {
	list := New[int, any]([]int{ 1, 2 })

	error, value := list.Last()
	if error != nil {
		t.Error("TestLast returned an error")		
	}
	if value != 2 {
		t.Error("TestLast did not return the expected value")
	}
}

func TestFirst(t *testing.T) {
	list := New[int, any]([]int{ 1, 2 })

	error, value := list.First()
	if error != nil {
		t.Error("TestGet returned an error")		
	}
	if value != 1 {
		t.Error("TestFirst did not return the expected value")
	}
}

func TestResultTypeSwap(t *testing.T) {
	list := New[int, any]([]int{ 1, 2 })

	type Number struct {
		Value int
	}

	newList := ResultTypeSwap[int, any, Number](list)

	reduced := newList.Reduce(func (accumulator Number, value int, index int) Number {
		return Number{ accumulator.Value + value }
	}, Number{ 22 })

	if reduced.Value != 25 {
		t.Error("TestResultTypeSwap did not return the expected empty value")
	}
}
