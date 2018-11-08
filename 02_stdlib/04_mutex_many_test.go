package stdlib_test

import (
	"sync"
	"testing"
)

var routinesMany = 100000

func BenchmarkMutexMany(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesMany)

	for ri := 0; ri < routinesMany; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incMutex()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkAtomicMany(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesMany)

	for ri := 0; ri < routinesMany; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incAtomic()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkChanMany(b *testing.B) {
	c := make(chan int64)

	go func() {
		for n := range c {
			testInt64 += n
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(routinesMany)

	for ri := 0; ri < routinesMany; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				c <- 1
			}
			wg.Done()
		}()
	}

	wg.Wait()

	close(c)
}

func BenchmarkChanBufferedMany(b *testing.B) {
	c := make(chan int64, routinesMany/10)

	wgc := sync.WaitGroup{}
	wgc.Add(1)

	go func() {
		for n := range c {
			testInt64 += n
		}
		wgc.Done()
	}()

	wg := sync.WaitGroup{}
	wg.Add(routinesMany)

	for ri := 0; ri < routinesMany; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				c <- 1
			}
			wg.Done()
		}()
	}

	wg.Wait()

	close(c)

	wgc.Wait()
}
func BenchmarkChanBufferedFullMany(b *testing.B) {
	c := make(chan int64, routinesMany*b.N)

	wgc := sync.WaitGroup{}
	wgc.Add(1)

	go func() {
		for n := range c {
			testInt64 += n
		}
		wgc.Done()
	}()

	wg := sync.WaitGroup{}
	wg.Add(routinesMany)

	for ri := 0; ri < routinesMany; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				c <- 1
			}
			wg.Done()
		}()
	}

	wg.Wait()

	close(c)

	wgc.Wait()
}
