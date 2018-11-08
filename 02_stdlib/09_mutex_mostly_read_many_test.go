package stdlib_test

import (
	"sync"
	"testing"
)

func BenchmarkMutexMostlyReadMany(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesMany)

	for ri := 0; ri < routinesMany/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incRWMutex()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesMany/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readRWMutex()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkAtomicMostlyReadMany(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(routinesMany)

	for ri := 0; ri < routinesMany/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				incAtomic()
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesMany/10*9; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				_ = readAtomic()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkChanMostlyReadMany(b *testing.B) {
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
	wg.Add(routinesMany)

	for ri := 0; ri < routinesMany/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				cw <- 1
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesMany/10*9; ri++ {
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

func BenchmarkChanBufferedMostlyReadMany(b *testing.B) {
	cw := make(chan int64, routinesMany/10)
	cr := make(chan int64, routinesMany/10)

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
	wg.Add(routinesMany)

	for ri := 0; ri < routinesMany/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				cw <- 1
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesMany/10*9; ri++ {
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

func BenchmarkChanBufferedFullMostlyReadMany(b *testing.B) {
	cw := make(chan int64, routinesMany)
	cr := make(chan int64, routinesMany)

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
	wg.Add(routinesMany)

	for ri := 0; ri < routinesMany/10; ri++ {
		go func() {
			for i := 0; i < b.N; i++ {
				cw <- 1
			}
			wg.Done()
		}()
	}

	for ri := 0; ri < routinesMany/10*9; ri++ {
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
