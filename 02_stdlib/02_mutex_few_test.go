package stdlib_test

import (
	"sync"
	"testing"
)

var routinesFew = 1000

func BenchmarkMutexFew(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesFew)

	for ri := 0; ri < routinesFew; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incMutex()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkAtomicFew(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesFew)

	for ri := 0; ri < routinesFew; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incAtomic()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkChanFew(b *testing.B) {
	c := make(chan int64)

	go func() {
		for n := range c {
			testInt64 += n
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(routinesFew)

	for ri := 0; ri < routinesFew; ri++ {
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

func BenchmarkChanBufferedFew(b *testing.B) {
	c := make(chan int64, routinesFew/10)

	wgc := sync.WaitGroup{}
	wgc.Add(1)

	go func() {
		for n := range c {
			testInt64 += n
		}
		wgc.Done()
	}()

	wg := sync.WaitGroup{}
	wg.Add(routinesFew)

	for ri := 0; ri < routinesFew; ri++ {
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

func BenchmarkChanBufferedFullFew(b *testing.B) {
	c := make(chan int64, routinesFew*b.N)

	wgc := sync.WaitGroup{}
	wgc.Add(1)

	go func() {
		for n := range c {
			testInt64 += n
		}
		wgc.Done()
	}()

	wg := sync.WaitGroup{}
	wg.Add(routinesFew)

	for ri := 0; ri < routinesFew; ri++ {
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
