package language_test

import "testing"

func BenchmarkDirectAnonymous(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			testInt64++
		}()
	}
}

func BenchmarkDeferAnonymous(b *testing.B) {
	for i := 0; i < b.N; i++ {
		incDeferAnonymous()
	}
}

func incDeferAnonymous() {
	defer func() { testInt64++ }()
}
