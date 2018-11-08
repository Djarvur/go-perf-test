package language_test

import "testing"

func BenchmarkDirectByPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		incDirectByPointer(&testInt64)
	}
}

func incDirectByPointer(n *int64) {
	*n++
}

func BenchmarkDeferByPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		incDeferByPointer(&testInt64)
	}
}

func incDeferByPointer(n *int64) {
	defer incDirectByPointer(n)
}
