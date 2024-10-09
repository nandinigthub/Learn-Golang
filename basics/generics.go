// introduced in 18 version  // we are using 21

package basics

import "fmt"

// data type defined
// func printSlice1(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// now this any or interface would take any data type
// func printSlice[T any](items []T) {
// 	// any ya interface dono use kar skte hai
// 	//func printSlice(T interface{})(items []T) {

// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// since this any or interface would take any dat type
// to restrict this

// func printSlice2[T int | string](items []T) {

// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//can also be used with structs

type stack[T any] struct {
	elements []T
}

// multiple generic types can also be used
func printSlice3[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

func Generic() {

	// nums := []int{1, 2, 3}
	// names := []string{"golang", "cpp", "java"}
	val := []bool{true, false}
	// printSlice(nums)
	// printSlice(names)
	// printSlice2(nums)
	// printSlice(val)
	// printSlice2(val) // error dega becoz bool not allowed
	printSlice3(val, "val")

	mystack := stack[string]{
		elements: []string{"aman", "rtyu"},
	}

	fmt.Println(mystack)

}
