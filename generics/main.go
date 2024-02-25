package main

// The `SumInts` function calculates the sum of all integer values in a given map.
// `Sum` that works for any `map` whose values are integers.

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func GenericSum[T comparable, V int64 | float64](m map[T]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	// ints := map[string]int64{
	ints := map[string]int64{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	floats := map[string]float64{
		"a": 1.1,
		"b": 2.2,
		"c": 3.3,
	}

	println(SumInts(ints))
	println(GenericSum[string, int64](ints))
	println(SumFloats(floats))
	println(GenericSum[string, float64](floats))
}
