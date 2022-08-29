package arrays

import (
	"fmt"
	"testing"
)

func TestChunk(t *testing.T) {
	source := []int{1, 4, 5, 9, 80, 76}

	r := Chunk(source, 0)

	fmt.Println(r)
}

func TestCountValues(t *testing.T) {
	source := []int{1, 4, 5, 9, 80, 76, 1, 5, 76}

	r := CountValues(source)

	fmt.Println(r)
}

func TestFill(t *testing.T) {
	r := Fill(10, "ee")

	fmt.Println(r)
}

func TestFilter(t *testing.T) {
	source := []int{1, 4, 5, 9, 80, 76, 1, 5, 76}

	r := Filter(source, func(item int, index int) bool {
		return item > 10
	})

	fmt.Println(r)
}

func TestFlip(t *testing.T) {
	source := []int{1, 4, 5, 9, 80, 76, 1, 5, 76}

	r := Flip(source)

	fmt.Println(r)
}

func TestPad(t *testing.T) {

	source := []int{1, 4, 6}

	r := Pad(source, -6, 800)

	fmt.Println(r)

}

func TestReverse(t *testing.T) {
	source := []int{1, 4, 6}

	r := Reverse(source)

	fmt.Println(r)
}

func TestSearch(t *testing.T) {
	source := []string{ "b", "a", "98"}

	r := Search(source, "a6")

	fmt.Println(r)
}

func TestProduct(t *testing.T) {
	source := []int{2, 4, 6}

	r := Product(source)

	fmt.Println(r)
}

func TestSum(t *testing.T) {
	source := []int{2, 4, 6}

	r := Sum(source)

	fmt.Println(r)
}

func TestUnique(t *testing.T) {
	source := []int{1, 4, 5, 9, 80, 76, 1, 5, 76, 8}

	r := Unique(source)

	fmt.Println(r)
}

func TestDiff(t *testing.T) {
	source := []int{1, 4, 5, 9, 80, 76, 1, 5, 76, 8}
	o := []int{1, 4, 5, 9, 80, 76, 1, 5, 76, }
	o2 := []int{1, 4, 5, 9, 80, 76, 1,  76, }

	r := Diff(source, o, o2)

	fmt.Println(r)
}

func TestIntersect(t *testing.T) {
	source := []int{1, 4, 5, 9, 80, 76, 1, 5, 76, 8}
	o := []int{1, 4, 5, 9, 80, 76, 1, 5, 76, }
	o2 := []int{1, 4,  9, 80, 76, 1,  76, }

	r := Intersect(source, o, o2)

	fmt.Println(r)
}