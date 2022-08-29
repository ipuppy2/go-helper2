package maps

// Column []
// PHP: array_column
func Column() {

}

// Combine []
// PHP: array_combine
func Combine [K comparable, V any] (keys []K, values []V) map[K]V {
	target := make(map[K]V)
	vLen := len(values)

	// 如果 value 小于 key
	if len(keys) > vLen {
		temp := make([]V, len(keys))
		copy(temp, values)
		values = temp
	}

	for i,k := range keys {
		target[k] = values[i]
	}
	return target
}

// KeyDiff []
// PHP: array_diff_key
func KeyDiff[K comparable, V any](source map[K]V, others ...map[K]V) map[K]V {
	return nil
}

// Diff []
// PHP: array_diff
func Diff[K comparable, V comparable](source map[K]V, others ...map[K]V) map[K]V {
	return nil
}

// FillKeys []
// PHP: array_fill_keys
func FillKeys[K comparable, V any](source []K, val V) map[K]V {
	target := make(map[K]V)
	for _, k := range source{
		target[k] = val
	}
	return target
}

// Flip []
// array_flip
func Flip [V, K comparable](source map[K]V) map[V]K {
	target := make(map[V]K)
	for k, v := range source {
		target[v] = k
	}
	return target
}

// KeyIntersect []
// PHP: array_intersect_key
func KeyIntersect [K comparable, V any] (source map[K]V, others ...map[K]V) map[K]V {
	return nil
}

// Intersect []
// PHP: array_intersect
func Intersect [K comparable, V any] (source map[K]V, others ...map[K]V) map[K]V {
	return nil
}

// Keys []
// PHP: array_keys
func Keys[K comparable, V any](source map[K]V) []K {
	target := make([]K, len(source))
	i := 0
	for k := range source{
		target[i] = k
		i ++
	}
	return target
}

// Values []
// PHP: array_values
func Values[K comparable, V any](source map[K]V) []V {
	target := make([]V, len(source))
	i := 0
	for _, v := range source{
		target[i] = v
		i ++
	}
	return target
}

// KeyAndValues [key and v]
func KeyAndValues[K comparable, V any](source map[K]V) ([]K, []V) {
	ks := make([]K, len(source))
	vs := make([]V, len(source))
	i := 0
	for k,v := range source{
		ks[i] = k
		vs[i] = v
		i++
	}
	return ks, vs
}

// Merge [直接修改 source]
// PHP:array_merge
func Merge[K comparable, V any](source map[K]V, others ...map[K]V) map[K]V {
	for _, item := range others{
		for k, v := range item{
			source[k] = v
		}
	}
	return source
}
