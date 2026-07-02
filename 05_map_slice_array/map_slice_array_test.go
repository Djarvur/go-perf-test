package map_slice_array

import (
	"testing"
)

func Benchmark_mapLoop(b *testing.B) {
	b.ReportAllocs()
	loopMap := make(map[string]bool, 9)
	loopMap["section1"] = false
	loopMap["section2"] = false
	loopMap["section3"] = false
	loopMap["section4"] = false
	loopMap["section5"] = false
	loopMap["section6"] = false
	loopMap["section7"] = false
	loopMap["section8"] = false
	loopMap["section9"] = false

	loopSlice := make(map[string]int, 9)
	loopSlice["section1"] = 0
	loopSlice["section2"] = 1
	loopSlice["section3"] = 2
	loopSlice["section4"] = 3
	loopSlice["section5"] = 4
	loopSlice["section6"] = 5
	loopSlice["section7"] = 6
	loopSlice["section8"] = 7
	loopSlice["section9"] = 8

	requestMap := map[string]string{
		"section1": "sdsection",
		"section2": "sdsection",
		"section3": "sdsection",
		"section4": "sdsection",
		"section5": "sdsection",
		"section6": "sdsection",
		"section7": "sdsection",
		"section8": "sdsection",
		"section9": "sdsection",
	}
	responseMap := map[string]string{
		"section1": "sdsection",
		"section2": "sdsection",
		"section3": "sdsection",
		"section4": "sdsection",
		"section5": "sdsection",
		"section6": "sdsection",
		"section7": "sdsection",
		"section8": "sdsection",
		"section9": "sdsection",
	}
	requestSlice := []string{"hss", "key", "hss", "key", "hss", "key", "hss", "key", "key"}
	responseSlice := []string{"hss", "key", "hss", "key", "hss", "key", "hss", "key", "key"}
	templateArray := [9]string{"section1", "section2", "section3", "section4", "section5", "section6", "section7", "section8", "section9"}
	requestArray := [10]string{"hss", "key", "hss", "key", "hss", "key", "hss", "key", "key", "key"}
	responseArray := [10]string{"hss", "key", "hss", "key", "hss", "key", "hss", "key", "key", "key"}
	b.Run("map based loop", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for key := range loopMap {
				_ = requestMap[key]
				_ = responseMap[key]
			}
		}
	})
	b.Run("slice based loop", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, index := range loopSlice {
				_ = requestSlice[index]
				_ = responseSlice[index]
			}
		}
	})
	b.Run("array based loop", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for index := range templateArray {
				shift := index + 1
				_ = requestArray[shift]
				_ = responseArray[shift]
			}
		}
	})
	// loop involves similar actions working with template maps/slices/arrays, so we have to understand the difference
	// 	Benchmark_mapLoop/map_based_loop-4         	 3000000	       415 ns/op	       0 B/op	       0 allocs/op
	//  Benchmark_mapLoop/slice_based_loop-4       	10000000	       140 ns/op	       0 B/op	       0 allocs/op
	//  Benchmark_mapLoop/array_based_loop-4       	500000000	         3.37 ns/op	       0 B/op	       0 allocs/op
}
