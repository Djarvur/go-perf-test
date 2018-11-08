package language_test

import (
	"testing"
)

//go:generate go run ./generator/generate_switch_data.go -template ./generator/switch_data_medium.template -file switch_data_medium_test.go -package language_test
//go:generate goimports -w switch_data_medium_test.go

func BenchmarkSwitchIntMedium(b *testing.B) {
	for i := 0; i < b.N; {
		for j := 0; j < len(randIntMedium) && i < b.N; j++ {
			switchIntMedium(randIntMedium[j])
			i++
		}
	}
}

func BenchmarkMapIntMedium(b *testing.B) {
	for i := 0; i < b.N; {
		for j := 0; j < len(randIntMedium) && i < b.N; j++ {
			testMapInt(mapIntMedium, randIntMedium[j])
			i++
		}
	}
}

func BenchmarkSliceIntMedium(b *testing.B) {
	for i := 0; i < b.N; {
		for j := 0; j < len(randIntMedium) && i < b.N; j++ {
			testSlice(sliceIntMedium, randIntMedium[j])
			i++
		}
	}
}

func BenchmarkSwitchStringMedium(b *testing.B) {
	for i := 0; i < b.N; {
		for j := 0; j < len(randIntMedium) && i < b.N; j++ {
			switchStringMedium(randStringMedium[j])
			i++
		}
	}
}

func BenchmarkMapStringMedium(b *testing.B) {
	for i := 0; i < b.N; {
		for j := 0; j < len(randStringMedium) && i < b.N; j++ {
			testMapString(mapStringMedium, randStringMedium[j])
			i++
		}
	}
}
