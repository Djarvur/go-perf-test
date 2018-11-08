package stdlib_test

import (
	"sync"
	"testing"
)

var routinesModerate = 10000

func BenchmarkMutexModerate(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesModerate)

	for ri := 0; ri < routinesModerate; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incMutex()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkAtomicModerate(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesModerate)

	for ri := 0; ri < routinesModerate; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incAtomic()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkChanModerate(b *testing.B) {
	c := make(chan int64)

	go func() {
		for n := range c {
			testInt64 += n
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(routinesModerate)

	for ri := 0; ri < routinesModerate; ri++ {
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

func BenchmarkChanBufferedModerate(b *testing.B) {
	c := make(chan int64, routinesModerate/10)

	wgc := sync.WaitGroup{}
	wgc.Add(1)

	go func() {
		for n := range c {
			testInt64 += n
		}
		wgc.Done()
	}()

	wg := sync.WaitGroup{}
	wg.Add(routinesModerate)

	for ri := 0; ri < routinesModerate; ri++ {
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
func BenchmarkChanBufferedFullModerate(b *testing.B) {
	c := make(chan int64, routinesModerate*b.N)

	wgc := sync.WaitGroup{}
	wgc.Add(1)

	go func() {
		for n := range c {
			testInt64 += n
		}
		wgc.Done()
	}()

	wg := sync.WaitGroup{}
	wg.Add(routinesModerate)

	for ri := 0; ri < routinesModerate; ri++ {
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
