package maps

import (
	"testing"
)

func TestMapDotMap(t *testing.T) {
	theMap := NewEmpty[string, string, string]()

	theMap.Set("hey", "dude")
	theMap.Set("sup", "brah")

	mappedList := theMap.Map(func(value string, key string, i int) string {
		return key + "_" + value
	})
	if error, value := mappedList.Get(0); error != nil || value != "hey_dude" {
		t.Error("unexpcted value from mappedList.Get(0) in TestMapDotMap, got " + value)
	}
}

func TestMapDotNew(t *testing.T) {
	nativeMap := make(map[string]string)
	nativeMap["sup"] = "brah"
	theMap := New[string, string, string](nativeMap)

	stringy := theMap.String()
	if stringy != "map[sup:brah]" {
		t.Error("unexpcted value in TestMapDotNew, got: " + stringy)
	}
}

func TestMapDotValues(t *testing.T) {
	theMap := NewEmpty[string, string, any]()
	theMap.Set("sup", "brah")
	stringy := theMap.Values().String()
	if stringy != "[brah]" {
		t.Error("unexpcted value in TestMapDotNew, got: " + stringy)
	}
}

func TestMapDotNewEmpty(t *testing.T) {
	theMap := NewEmpty[string, string, string]()
	stringy := theMap.String()
	if stringy != "map[]" {
		t.Error("unexpected value in TestMapDotNewEmpty, got: " + stringy)
	}
}

func TestMapDotReduce(t *testing.T) {
	theMap := NewEmpty[string, string, string]()

	theMap.Set("hey", "dude")
	theMap.Set("sup", "brah")

	initial := "When I meet someone new, I always say: "
	message := theMap.Reduce(func(accumulator string, value string, key string, i int) string {
		return accumulator + key + " " + value + " "
	}, initial)
	expected := initial + "hey dude sup brah "
	if message != expected {
		t.Error("oh noes, got" + message + "in TestMapDotReduce")
	}
}

func TestMapDotString(t *testing.T) {
	theMap := NewEmpty[string, string, string]()

	theMap.Set("hey", "dude")
	theMap.Set("sup", "brah")

	stringy := theMap.String()

	if stringy != "map[hey:dude sup:brah]" {
		t.Error("oh noes, got" + stringy + "in TestMapDotString")
	}
}