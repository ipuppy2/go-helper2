package arrays

import (
	"math"
	"puppy/types"
)

// Chunk [chunk array]
// PHP: array_chunk()
func Chunk[T any](source []T, size int) [][]T {
	if size == 0 {
		target := make([][]T, 1)
		target[0] = source
		return target
	}
	chunk := make([][]T, 0, int(math.Ceil(float64(len(source)/size))))
	length := len(source)
	for index := 0; index < length; index += size {
		end := size + index
		if end > length {
			end = length
		}
		chunk = append(chunk, source[index:end])
	}

	return chunk
}

// CountValues []
// PHP: array_count_values
func CountValues[V comparable](source []V) map[V]int {
	cv := make(map[V]int)
	for _, it := range source{
		cv[it] ++
	}
	return cv
}

// Diff []
// PHP: array_diff
func Diff[V comparable](source []V, others ...[]V) []V {

	if len(others) == 0 {
		target := make([]V, len(source))
		copy(target, source)
		return target
	}

	has := make(map[V]struct{})
	target := make([]V, 0)
	for _, v := range source{

		if _, ex := has[v]; ex {
			continue
		}

		found := false
		for _, list := range others {
			for _, ov := range list{
				has[ov] = struct{}{}
				if ov == v {
					found = true
					break
				}
			}
			if found {
				break
			}
		}

		if !found {
			target = append(target, v)
		}

	}

	return target
}

// Fill []
// PHP: array_fill
func Fill[V any](size int, val V) []V {
	target := make([]V, size)
	for i:=0; i<size; i++ {
		target[i] = val
	}
	return target
}

// Filter []
// PHP: array_filter
func Filter [V any] (source []V, f func(item V, index int) bool) []V {
	target := make([]V, 0)
	for i, val := range source{
		if f(val, i) {
			target = append(target, val)
		}
	}
	return target
}

// IsEmpty []
func IsEmpty [V any] (source []V) bool  {
	return len(source) == 0
}

// NotEmpty []
func NotEmpty [V any](source []V) bool {
	return !IsEmpty(source)
}

// Flip []
// array_flip
func Flip [V comparable](source []V) map[V]int {
	target := make(map[V]int)
	for i, v := range source {
		target[v] = i
	}
	return target
}

// Intersect []
// PHP: array_intersect
func Intersect [V comparable] (source []V, others ...[]V) []V {
	capSize := len(others)
	if 0 == capSize {
		return make([]V, 0)
	}

	target := make([]V, 0)

	pool := make(map[V]int)


	for _, val := range source {

		times := pool[val]
		if times == capSize {
			// 都存在
			target = append(target, val)
			continue
		}

		found := true
		for _, list := range others {

			subFound := false
			for _, ov := range list{
				pool[ov] ++
				if ov == val {
					subFound = true
					break
				}
			}
			found = subFound && found
			if !found {
				break
			}
		}

		if found {
			// 都存在
			target = append(target, val)
		}

	}

	return target

}

// Pad []
// PHP: array_pad
// 如果 newLength < 0, 则从填充到左边，否则，填充到右边
// 如果  newLength 的 abs 小于或等于 原长度，则无任何改变

func Pad [V any] (source []V, newLength int, padV V) []V {
	length := len(source)
	absLen := newLength
	if newLength < 0 {
		absLen = newLength * -1
	}

	if absLen <= length {
		return source
	}

	diffLen := absLen - length

	padValues := make([]V, diffLen)
	for i := 0; i<diffLen; i++ {
		padValues[i] = padV
	}

	if newLength < 0 {
		return append(padValues, source...)
	}

	return append(source, padValues...)
}

// Reverse [直接修改原数组]
// PHP: array_reverse
func Reverse [V any] (source []V) []V  {
	inputLen := len(source)
	inputMid := inputLen / 2

	for i := 0; i < inputMid; i++ {
		j := inputLen - i - 1

		source[i], source[j] = source[j], source[i]
	}
	return source
}

// Search
// PHP: array_search
func Search [V comparable] (source []V, find V) int  {
	length := len(source)
	for index := 0; index<length; index++ {
		if source[index] == find {
			return index
		}
	}

	return -1
}

/*func SearchSorted[V comparable] (source []V, find V) int  {
	index := sort.Search(len(source), func(i int) bool {
		return source[i] >= find
	})

	return index
}*/

// Product []
// PHP: array_product
func Product [V types.Calculable] (source []V) V {

	length := len(source)
	if length == 0 {
		return 0
	}

	target := source[0]
	for i:=1; i<length; i++ {
		target *= source[i]
	}
	return target
}

// Sum
// PHP: array_sum
func Sum [V types.Calculable] (source []V) V {

	length := len(source)
	if length == 0 {
		return 0
	}

	target := source[0]
	for i:=1; i<length; i++ {
		target += source[i]
	}
	return target
}

// Unique
// PHP:array_unique
func Unique[V comparable](source []V) []V {
	if len(source) == 0 {
		return make([]V, 0)
	}

	fm := make(map[V]int)
	for i, v := range source {
		if _, h := fm[v]; h {
			continue
		}
		fm[v] = i
	}
	filter := make([]V, len(fm))
	index := 0
	for _, i := range fm{
		filter[index] = source[i]
		index++
	}
	return filter
}


// array_change_key_case
// array_column
// array_combine
// array_count_values
// array_diff_assoc
// array_diff_key
// array_diff_uassoc
// array_diff_ukey
// array_diff
// array_fill_keys
// array_fill
// array_filter
// array_flip
// array_intersect_assoc
// array_intersect_key
// array_intersect_uassoc
// array_intersect_ukey
// array_intersect
// array_is_list
// array_key_exists
// array_key_first
// array_key_last
// array_keys
// array_map
// array_merge_recursive
// array_merge
// array_multisort
// array_pad
// array_pop
// array_push
// array_rand
// array_reduce
// array_replace_recursive
// array_replace
// array_reverse
// array_shift
// array_slice
// array_splice
// array_udiff_assoc
// array_udiff_uassoc
// array_udiff
// array_uintersect_assoc
// array_uintersect_uassoc
// array_uintersect
// array_unshift
// array_values
// array_walk_recursive
// array_walk
// array
// arsort
// asort
// compact
// count
// current
// end
// extract
// in_array
// key_exists
// key
// krsort
// ksort
// list
// natcasesort
// natsort
// next
// pos
// prev
// range
// reset
// rsort
// shuffle
// sizeof
// sort
// uasort
// uksort
// usort
