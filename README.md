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
array := []int { 1, 2 }
list := New[int, int](array)
added := list.Map(func(val int, index int) int {
	return val + 1
}).Reduce(func(accumulator int, val int, index int) int {
	return accumulator + val
}, 0)
fmt.Println(added) // -> 5
```

## Map (still unstable)

```
import "github.com/neurocollective/go_chainable/lists"

nativeMap := map[string]string {
	"hey": "dude",
	"sup": "brah",
}
theMap := maps.New[string, string, string](nativeMap)

mapped := theMap.Map(func(value string, key string) string {
	return key + "_" + value
})

fmt.Println(mapped.String()) // -> [hey_dude sup_brah]
```

## Current Test Coverage

```
ok  	github.com/neurocollective/go_chainable/lists	0.002s	coverage: 53.6% of statements
ok  	github.com/neurocollective/go_chainable/maps	0.001s	coverage: 40.0% of statements
```