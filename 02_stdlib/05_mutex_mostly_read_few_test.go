package stdlib_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

var lockRW = sync.RWMutex{}

func BenchmarkMutexMostlyReadFew(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesFew)

	for ri := 0; ri < routinesFew/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incRWMutex()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesFew/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readRWMutex()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func incRWMutex() {
	lockRW.Lock()
	testInt64++
	lockRW.Unlock()
}

func readRWMutex() int64 {
	lockRW.RLock()
	res := testInt64
	lockRW.RUnlock()
	return res
}

func BenchmarkAtomicMostlyReadFew(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesFew)

	for ri := 0; ri < routinesFew/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incAtomic()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesFew/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readAtomic()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func readAtomic() int64 {
	return atomic.LoadInt64(&testInt64)
}

func BenchmarkChanMostlyReadFew(b *testing.B) {
	cw := make(chan int64)
	cr := make(chan int64)

	go func() {
		for {
			select {
			case n, ok := <-cw:
				if !ok {
					return
				}
				testInt64 += n
			case cr <- testInt64:
			}
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(routinesFew)

	for ri := 0; ri < routinesFew/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				cw <- 1
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesFew/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = <-cr
			}
			wg.Done()
		}()
	}

	wg.Wait()

	close(cw)
}

func BenchmarkChanBufferedMostlyReadFew(b *testing.B) {
	cw := make(chan int64, routinesFew/10)
	cr := make(chan int64, routinesFew/10)

	wgc := sync.WaitGroup{}
	wgc.Add(1)

	go func() {
		for {
			select {
			case n, ok := <-cw:
				if !ok {
					wgc.Done()
					return
				}
				testInt64 += n
			case cr <- testInt64:
			}
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(routinesFew)

	for ri := 0; ri < routinesFew/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				cw <- 1
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesFew/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = <-cr
			}
			wg.Done()
		}()
	}

	wg.Wait()

	close(cw)

	wgc.Wait()
}

func BenchmarkChanBufferedFullMostlyReadFew(b *testing.B) {
	cw := make(chan int64, routinesFew)
	cr := make(chan int64, routinesFew)

	wgc := sync.WaitGroup{}
	wgc.Add(1)

	go func() {
		for {
			select {
			case n, ok := <-cw:
				if !ok {
					wgc.Done()
					return
				}
				testInt64 += n
			case cr <- testInt64:
			}
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(routinesFew)

	for ri := 0; ri < routinesFew/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				cw <- 1
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesFew/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = <-cr
			}
			wg.Done()
		}()
	}

	wg.Wait()

	close(cw)

	wgc.Wait()
}
