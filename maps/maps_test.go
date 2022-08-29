package maps

import (
	"fmt"
	"testing"
)

func TestCombine(t *testing.T) {
	ks := []string{"a", "b", "3"}
	vs := []struct{}{{}}

	r := Combine(ks, vs)
	fmt.Println(r)
}

func TestFillKeys(t *testing.T) {
	ks := []string{"a", "b", "3"}

	r := FillKeys(ks, "44")
	fmt.Println(r)
}

func TestFlip(t *testing.T) {
	ks := make(map[string]int)
	ks["a"] = 4
	ks["b"] = 34
	ks["c"] = 49
	r := Flip(ks)
	fmt.Println(r)
}

func TestKeys(t *testing.T) {
	ks := make(map[string]int)
	ks["a"] = 4
	ks["b"] = 34
	ks["c"] = 49
	r := Keys(ks)
	fmt.Println(r)
}

func TestValues(t *testing.T) {
	ks := make(map[string]int)
	ks["a"] = 4
	ks["b"] = 34
	ks["c"] = 49
	r := Values(ks)
	fmt.Println(r)
}

func TestKeyAndValues(t *testing.T) {
	ks := make(map[string]int)
	ks["a"] = 4
	ks["b"] = 34
	ks["c"] = 49
	k, v := KeyAndValues(ks)
	fmt.Println(k, v)
}

func TestMerge(t *testing.T) {
	ks := make(map[string]int)
	ks["a"] = 4
	ks["b"] = 34
	ks["c"] = 49

	ks2 := make(map[string]int)
	ks2["a-"] = 4
	ks2["b"] = 34000
	ks2["c-"] = 49

	r := Merge(ks, nil)
	fmt.Println(r)
}