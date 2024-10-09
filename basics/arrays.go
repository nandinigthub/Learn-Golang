package basics

import (
	"fmt"
	"slices"
)

func Array1() {
	// var nums [4]int
	// nums[0] = 2
	// nums[3] = 4

	// fmt.Println(nums)

	// var s [6]string
	// s[0] = "sneha"
	// s[1] = "ruby"
	// s[2] = "nandini"
	// s[3] = "raman"
	// fmt.Println(s)
	// fmt.Println(s[3])
	// fmt.Println(len(s))

	// // single line declaration
	// arr := [3]int{1, 2, 3}
	// fmt.Println(arr)
	// // 2d array
	// arr2 := [2][2]int{{1, 2}, {3, 4}}
	// fmt.Println(arr2)

	// slices - variable array
	// var slice []int
	// fmt.Println(slice == nil)
	// fmt.Println(len(slice))

	// Creating an array whose size
	// is represented by the ellipsis
	// my_array := [...]int{100, 200, 300, 400, 500}
	// fmt.Println("Original array:", my_array)

	// // Creating a new variable
	// // and initialize with my_array

	// new_array := my_array
	// fmt.Println("coppied array:", new_array)

	// // declare using make func
	// var arr = make([]int, 5)
	// var new_arr = make([]int, len(arr))

	// // copy(des,source)
	// copy(new_arr, arr)
	// fmt.Println(new_arr)

	// slice operator
	// just like python
	var nums1 = []int{1, 2, 3}
	var nums2 = []int{1, 2, 3}
	fmt.Println(slices.Equal(nums1, nums2))
	fmt.Println(nums1[:2])
}
