package basics

import "fmt"

// func add(a, b int) int {
// 	return a + b
// }

// type of return ; order of type matter
func getLang() (string, int, string) {
	return "go", 20, "c"
}

// function inside a func
// fn func(a int) int
//This indicates that fn is a function
// that takes an integer a as an argument and returns an integer.

func doublefunc(fn func(a int) int) int {
	return fn(2) + 5
}

// check palindrome func
func chkpalindrome(s string) bool {
	i := 0
	j := len(s)
	for i <= j {
		if s[i] != s[j-1] {
			return false
		}
		i++
		j--

	}
	return true
}

func Fuct() {
	// res := add(3, 5)
	// fmt.Println(res)

	fmt.Println(getLang())

	fn := func(a int) int {
		return 2 * a
	}

	fmt.Println(doublefunc(fn))
	fmt.Println(chkpalindrome("naman"))
	fmt.Println(chkpalindrome("namaya"))

}
