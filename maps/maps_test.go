package maps

import (
	"testing"
	"github.com/neurocollective/go_chainable/lists"
)

func TestMapDotMap(t *testing.T) {
	mapperTwo := func(value string, key string, index int, nativeMap *map[string]string) string {
		return key + "_" + value
	}	

	nativeMap := map[string]string {
		"hey": "dude",
		"sup": "brah",
	}
	array := []string { "hey", "sup" }
	theMap := Map[string, string, string]{
		&nativeMap,
		&lists.List[string, string]{ &array, "" },
	}

	mappedTwo := theMap.Map(mapperTwo)

	if error, value := mappedTwo.Get(0); error != nil || value != "hey_dude" {
		t.Error("oh noes")
	}
}