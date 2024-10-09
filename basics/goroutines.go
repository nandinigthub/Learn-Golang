// helps to run simultaneously ithout blocking
// uses lightweight thread
// simply use go keyword

package basics

import (
	"fmt"
	"sync"
)

func loops(a int, w *sync.WaitGroup) {
	// defer runs after func completes
	defer w.Done()
	fmt.Println("task", a)
}

func Goroutines() {

	// WaitGroup
	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go loops(i, &wg)
		// also write inline goroutine function
		// go func(i int) {
		// 	fmt.Println(i)
		// }(i)

	}

	wg.Wait()

	// time.Sleep(time.Second * 2)
	// we cant use time sleep in real world we dont know the time beforehand
	// there fore we use WaitGroup

}
