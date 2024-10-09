package basics

import "fmt"

func Defers() {

	// a := func() {
	// 	i := 0
	// 	defer fmt.Println("prints i", i) // value purani hi store hogi defer m but time sleep k baad execute hoga before returning val
	// 	i++
	// 	time.Sleep(time.Second)
	// 	defer fmt.Println("after sleep prints i", i)
	// }

	// a()

	// Deferred function calls are executed in Last In First Out order after the surrounding function returns.
	// This function prints “3210”:

	// b := func() {
	// 	for i := 0; i < 4; i++ { // basically defer se loop ulta chalne lagta hai, hahahah xd
	// 		defer fmt.Println("prints value of i after evvery loop execution", i)
	// 	}
	// }
	// b()

	// Deferred functions may read and assign to the returning function’s named return values.
	c := func(i int) int {
		defer func() {
			i++
			fmt.Println("Deferred i:", i) // This will print the incremented value of i
		}()
		return 1
	}

	result := c(5) // Passing 5 as the argument to c
	fmt.Println("Return value:", result)
}
