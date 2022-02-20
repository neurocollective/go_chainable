package main

import (
	"fmt"
	"neurocollective.io/godash/lists"
	"neurocollective.io/godash/maps"
)

// func mapperTwo[s string](value string, key string, nativeMap *map[string]string) string {
// 	return key + "_" + value
// }

	func mapperTwo(value string, key string, nativeMap *map[string]string) string {
		return key + "_" + value
	}

func main() {
	fmt.Println("sup")

	arr := []int { 0, 1, 2, 3, 4 }
	list := lists.List[int]{ &arr }

	fmt.Println(*list.Array)

	mapper := func (value int, index int, array *[]int) int {
		return value + 1
	}

	mapped := list.Map(mapper)

	fmt.Println(*mapped)

	nativeMap := map[string]string {
		"hey": "dude",
		"sup": "brah",
	}
	array := []string { "hey", "sup" }
	theMap := maps.Map[string, string]{
		&nativeMap,
		&lists.List[string]{ &array },
	}

	mappedTwo := maps.FunctionalMapping[string, string, string](&theMap, mapperTwo)

	fmt.Println(*mappedTwo)

}