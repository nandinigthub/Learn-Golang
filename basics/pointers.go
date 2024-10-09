package basics

import "fmt"

//by value
// func namechange(numx int) {
// 	numx = 5
// 	fmt.Println("before", numx)
// }

// by reference
func namechange2(nums *int) {
	*nums = 5
	fmt.Println("before2", nums)
}

func Point() {

	num := 1
	// namechange(num)
	// fmt.Println("after", num)
	// address of num
	fmt.Println("address", &num)

	namechange2(&num)
	fmt.Println("after2", &num)

}
