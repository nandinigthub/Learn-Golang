package basics

import "fmt"

// closure - ek func k andar doosra func hota hai to uska variable use karta hai
// same as javascript - remembers prev variable val

func Counter() func() int {
	var cnt int = 1
	// anonymous function - func with no name
	return func() int {
		for cnt < 10 {
			cnt += 1
			fmt.Println(cnt)
		}
		return cnt
	}
}

// func Xyz() {
// 	fmt.Println("running xyz fnc")
// }

func Closures() {

	increment := Counter()
	fmt.Println(increment())

}
