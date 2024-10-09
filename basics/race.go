package basics

import (
	"fmt"
	"sync"
)

func Race() {

	fmt.Println("race conditions running....")

	// checking race condition - go run --race .
	// Found 4 data race(s) - exit status 66
	// to overcome this we use Mutex
	mut := &sync.Mutex{}
	Rmut := &sync.RWMutex{}

	wg := &sync.WaitGroup{} // passed by pointer
	var score = []int{0}

	wg.Add(4) // jitne func honge utne ek sath
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("1st race")
		mut.Lock()
		score = append(score, 1)
		mut.Lock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1) // alag alag bhi add use karskte hai doent make any difference
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("2st race")
		mut.Lock()
		score = append(score, 2)
		mut.Lock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("3st race")
		mut.Lock()
		score = append(score, 3)
		mut.Lock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		Rmut.RLock()
		fmt.Print(score)
		Rmut.RUnlock()
		wg.Done()
	}(wg, Rmut)

	//wg.Add(1)

	wg.Wait()
	fmt.Println(score)

}
