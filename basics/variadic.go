// variable functions are variadic functions

package basics

import "fmt"

// variable size - phele dot
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}
	return total
}

func Varii() {
	res := sum(1, 2, 3, 4, 5)
	fmt.Println(res)

	// slice - baad m dots
	arr := []int{1, 23, 4}
	ans := sum(arr...)
	fmt.Println(ans)
}
