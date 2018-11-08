package language_test

import "testing"

var testInt64 int64

func BenchmarkDirect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		incDirect()
	}
}

func incDirect() {
	testInt64++
}

func BenchmarkDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		incDefer()
	}
}

func incDefer() {
	defer incDirect()
}

func BenchmarkDirectNoInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		incDirectNoInline()
	}
}

//go:noinline
func incDirectNoInline() {
	testInt64++
}

func BenchmarkDeferNoInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		incDeferNoInline()
	}
}

//go:noinline
func incDeferNoInline() {
	defer incDirect()
}
