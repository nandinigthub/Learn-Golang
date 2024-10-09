package basics

import (
	"fmt"
)

func Ranges() {

	// nums := []int{1, 2, 3, 4, 5}
	// for _, val := range nums {
	// 	fmt.Println(val)
	// }

	// nums2 := [2][2]int{{1, 2}, {3, 4}}
	// for key, val := range nums2 {
	// 	fmt.Println(key, val)
	// }

	// unicode point rune
	// starting byte of rune
	// 300 ->1 byte, 2 byte

	for i, c := range "golang" {
		// c basically gives ascii value
		// fmt.Println(i, string(c))
		// fmt.Println(i, (c))
		fmt.Println(i, string(c), (c))

	}
}
