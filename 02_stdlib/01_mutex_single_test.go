package stdlib_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

var (
	testInt64 int64
	lock      sync.Mutex
)

func BenchmarkDirect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		incDirect()
	}
}

func incDirect() {
	testInt64++
}

func BenchmarkMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		incMutex()
	}
}

func incMutex() {
	lock.Lock()
	testInt64++
	lock.Unlock()
}

func BenchmarkAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		incAtomic()
	}
}

func incAtomic() {
	atomic.AddInt64(&testInt64, 1)
}

func BenchmarkChan(b *testing.B) {
	c := make(chan int64)

	go func() {
		for n := range c {
			testInt64 += n
		}
	}()

	for i := 0; i < b.N; i++ {
		c <- 1
	}

	close(c)
}
