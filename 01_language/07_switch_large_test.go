package language_test

import (
	"testing"
)

//go:generate go run ./generator/generate_switch_data.go -template ./generator/switch_data_large.template -file switch_data_large_test.go -package language_test
//go:generate goimports -w switch_data_large_test.go

func BenchmarkSwitchIntLarge(b *testing.B) {
	for i := 0; i < b.N; {
		for j := 0; j < len(randIntLarge) && i < b.N; j++ {
			switchIntLarge(randIntLarge[j])
			i++
		}
	}
}

func BenchmarkMapIntLarge(b *testing.B) {
	for i := 0; i < b.N; {
		for j := 0; j < len(randIntLarge) && i < b.N; j++ {
			testMapInt(mapIntLarge, randIntLarge[j])
			i++
		}
	}
}

func BenchmarkSliceIntLarge(b *testing.B) {
	for i := 0; i < b.N; {
		for j := 0; j < len(randIntLarge) && i < b.N; j++ {
			testSlice(sliceIntLarge, randIntLarge[j])
			i++
		}
	}
}

func BenchmarkSwitchStringLarge(b *testing.B) {
	for i := 0; i < b.N; {
		for j := 0; j < len(randIntLarge) && i < b.N; j++ {
			switchStringLarge(randStringLarge[j])
			i++
		}
	}
}

func BenchmarkMapStringLarge(b *testing.B) {
	for i := 0; i < b.N; {
		for j := 0; j < len(randStringLarge) && i < b.N; j++ {
			testMapString(mapStringLarge, randStringLarge[j])
			i++
		}
	}
}
