package language_test

import (
	"fmt"
	"testing"
)

//go:generate go run ./generator/generate_switch_data.go -template ./generator/switch_data_small.template -file switch_data_small_test.go -package language_test
//go:generate goimports -w switch_data_small_test.go

func BenchmarkSwitchIntSmall(b *testing.B) {
	for i := 0; i < b.N; {
		for j := 0; j < len(randIntSmall) && i < b.N; j++ {
			switchIntSmall(randIntSmall[j])
			i++
		}
	}
}

func BenchmarkMapIntSmall(b *testing.B) {
	for i := 0; i < b.N; {
		for j := 0; j < len(randIntSmall) && i < b.N; j++ {
			testMapInt(mapIntSmall, randIntSmall[j])
			i++
		}
	}
}

func testMapInt(m map[int]func(), p int) {
	f := m[p]
	if f == nil {
		panic(fmt.Errorf("Unexpected parameter: %v", p))
	}
	f()
}
func BenchmarkSliceIntSmall(b *testing.B) {
	for i := 0; i < b.N; {
		for j := 0; j < len(randIntSmall) && i < b.N; j++ {
			testSlice(sliceIntSmall, randIntSmall[j])
			i++
		}
	}
}

func testSlice(m []func(), p int) {
	f := m[p]
	f()
}

func BenchmarkSwitchStringSmall(b *testing.B) {
	for i := 0; i < b.N; {
		for j := 0; j < len(randIntSmall) && i < b.N; j++ {
			switchStringSmall(randStringSmall[j])
			i++
		}
	}
}

func BenchmarkMapStringSmall(b *testing.B) {
	for i := 0; i < b.N; {
		for j := 0; j < len(randStringSmall) && i < b.N; j++ {
			testMapString(mapStringSmall, randStringSmall[j])
			i++
		}
	}
}

func testMapString(m map[string]func(), p string) {
	f := m[p]
	if f == nil {
		panic(fmt.Errorf("Unexpected parameter: %v", p))
	}
	f()
}
