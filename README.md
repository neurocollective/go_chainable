# go_chainable

`NOTE: Requires go1.18`

Go-Chainable is a library using generics to mimic the functional `.map(x -> y).reduce(x -> y, z)` patterns of javascript. `List` is also inspired by Java's `ArrayList`.

## List

### Construct a List

```
array := []string { "hello", "world" }
list := lists.New[string, any](array)
```

or get an empty list

```
list := lists.NewEmpty[string, any]()
```

#### First type parameter

First type parameter is the type of the List's values.

#### "Result" type parameter

The second type parameter for a list is needed only for when calling `list.Reduce()`. If you won't need to call `.Reduce()` then set the second type argument as `any` and forget it.

### Map Over A List

```
import "github.com/neurocollective/go_chainable/lists"

array := []string { 1, 2 }
list := lists.New[int, any](array)
doubled := list.Map(func(val string, index int) int {
	return val * 2
})
fmt.Println(doubled.String()) // -> [2 4]
```

### Filter A List

```
array := []int { 1, 2 }
list := lists.New[int, any](array)
newList := list.Filter(func(val int, index int) bool {
	return val < 2
})
error, size := newList.Size()
fmt.Println(size) // -> 1
```

### Reduce A List

```
import "github.com/neurocollective/go_chainable/lists"

array := []int { 1, 2 }
list := lists.New[int, int](array) // second type passed to New is the "result" type used by reducer as return type
added := list.Reduce(func(accumulator int, val int, index int) int {
	return accumulator + val
}, 0)
fmt.Println(added) // -> 3
```

### Chain Operations Over A List

```
array := []int { 1, 2, 3 }
list := lists.New[int, int](array)
added := list.Map(func(val int, index int) int {
	return val + 1
}).Filter(func(val int, index int) bool {
	return val < 4
}).Reduce(func(accumulator int, val int, index int) int {
	return accumulator + val
}, 0)
fmt.Println(added) // -> 5
```

## Map

### `.Map` over key/value pairs

```
import "github.com/neurocollective/go_chainable/maps"

theMap := maps.NewEmpty[string, string, string]()

theMap.Set("hey", "dude")
theMap.Set("sup", "brah")

mappedList := theMap.Map(func(value string, key string, i int) string {
	return key + "_" + value
})

fmt.Println(mapped.String()) // -> [hey_dude sup_brah]
```

## `.Reduce` over key/value pairs

```
	theMap := NewEmpty[string, string, string]()

	theMap.Set("hey", "dude")
	theMap.Set("sup", "brah")

	initial := "When I meet someone new, I always say: "
	message := theMap.Reduce(func(accumulator string, value string, key string, i int) string {
		return accumulator + key + " " + value + " "
	}, initial)
```

## Current Test Coverage

```
ok  	github.com/neurocollective/go_chainable/lists	0.002s	coverage: 80.4% of statements
ok  	github.com/neurocollective/go_chainable/maps	0.001s	coverage: 97.3% of statements
```

## Methods

### List Methods

`.Raw()`

Returns a raw underlying `[]T` 

`.RawPointer()`

Returns the raw Pointer to the underlying `[]T`

`.Map(mapper func(value T, index int) T) *List[T, R]`

Calls the `mapper` function for each element, passing in `value` and `int`. The underlying array is changed to a new array, with each element being the returned value from `mapper`. 

`.MapFull(mapper func(value T, index int, array *[]T) T) *List[T, R]`

Same as `.Map` but the `mapper` function takes a pointer to the underlying array as an additional argument.

`.Reduce(reducer func(accumulator R, value T, index int) R, initial R) R)`

Calls the `reducer` function for each element, passing in `accumulator` `value` and `int`. On the first call to `reducer` the `accumulator` value is the `initial` value. But each subsequent call receives an `accumulator` that is the returned `R` value from the previous call to `reducer`.  

### List Functions

### Map Methods

### Map Functions