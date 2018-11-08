package passingparams_test

import (
	"math/rand"
	"testing"
)

//go:generate go run ./generator/generate_nesting_types.go -file passing_types_test.go -package passingparams_test -nesting 10
//go:generate goimports -w passing_types_test.go

var (
	smallByValue    = make([]TV002, 100)
	smallByPointer  = make([]*TP002, 100)
	mediumByValue   = make([]TV005, 100)
	mediumByPointer = make([]*TP005, 100)
	largeByValue    = make([]TV007, 100)
	largeByPointer  = make([]*TP007, 100)
)

func BenchmarkCreateSmallByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for ri := 0; ri < len(smallByValue); ri++ {
			smallByValue[ri] = NewTV002()
		}
	}
}

func BenchmarkCreateSmallByPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for ri := 0; ri < len(smallByValue); ri++ {
			smallByPointer[ri] = NewTP002()
		}
	}
}

func BenchmarkCreateMediumByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for ri := 0; ri < len(smallByValue); ri++ {
			mediumByValue[ri] = NewTV005()
		}
	}
}

func BenchmarkCreateMediumByPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for ri := 0; ri < len(smallByValue); ri++ {
			mediumByPointer[ri] = NewTP005()
		}
	}
}

func BenchmarkCreateLargeByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for ri := 0; ri < len(smallByValue); ri++ {
			largeByValue[ri] = NewTV007()
		}
	}
}

func BenchmarkCreateLargeByPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for ri := 0; ri < len(smallByValue); ri++ {
			largeByPointer[ri] = NewTP007()
		}
	}
}

func BenchmarkReadSmallByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = smallByValue[rand.Intn(len(smallByValue))]
	}
}

func BenchmarkReadSmallByPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = smallByPointer[rand.Intn(len(smallByPointer))]
	}
}

func BenchmarkReadMediumByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = mediumByValue[rand.Intn(len(mediumByValue))]
	}
}

func BenchmarkReadMediumByPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = mediumByPointer[rand.Intn(len(mediumByPointer))]
	}
}

func BenchmarkReadLargeByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = largeByValue[rand.Intn(len(largeByValue))]
	}
}

func BenchmarkReadLargeByPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = largeByPointer[rand.Intn(len(largeByPointer))]
	}
}
