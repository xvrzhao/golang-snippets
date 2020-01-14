package race

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// IntRace1 demonstrates the one of non-concurrency-safe scenarios of int value.
// Run the following command under the project root directory to test race:
//   $ go test -race -run=IntRace1 -v -count=1 ./race
func IntRace1() {
	var i int
	done := make(chan bool)
	go func() {
		i = 9 // unsafe
		done <- true
	}()
	i = 7 // unsafe
	<-done
	fmt.Println(i)
}

// IntRace2 demonstrates the one of non-concurrency-safe scenarios of int value.
// Run the following command under the project root directory to test race:
//   $ go test -race -run=IntRace2 -v -count=1 ./race
func IntRace2() {
	var num int
	wg := new(sync.WaitGroup)
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			num++ // unsafe
		}()
	}
	wg.Wait()
	fmt.Println(num) // num will be less than 1000
}

// IntAtomic demonstrates the one of concurrency-safe scenarios of int value.
// Run the following command under the project root directory to test race:
//   $ go test -race -run=IntAtomic -v -count=1 ./race
func IntAtomic() {
	var i int64
	done := make(chan bool)
	go func() {
		atomic.StoreInt64(&i, 9)
		done <- true
	}()
	atomic.StoreInt64(&i, 7)
	<-done
	fmt.Println(i)
}
