package basics

import (
	"fmt"
)

func Loop() {

	// for i := 0; i <= 5; i++ {
	// 	if i == 2 {
	// 		continue
	// 	}

	// 	if i == 4 {
	// 		break
	// 	}

	// 	fmt.Println(i)
	// }

	// range starts from 0 - n-1
	// for i := range 11 {
	// 	fmt.Println(i)
	// }

	// if else

	// age := 17

	// if age >= 18 {
	// 	fmt.Println("adult")
	// } else {
	// 	fmt.Println("not adult")
	// }

	// // directly declare variable
	// if age := 45; age >= 18 {
	// 	fmt.Println("adult")
	// } else {
	// 	fmt.Println("not adult")
	// }

	// switch

	// i := 13

	// switch i {
	// case 1:
	// 	fmt.Println("1")
	// case 2:
	// 	fmt.Println("2")
	// case 3:
	// 	fmt.Println("3")
	// case 4:
	// 	fmt.Println("4")
	// default:
	// 	fmt.Println("not present")

	// }

	// // multiple switch
	// // current day
	// switch time.Now().Weekday() {
	// case time.Sunday, time.Saturday:
	// 	fmt.Println("weekend")
	// default:
	// 	fmt.Println("workkk")

	// }

	// interface()

	value := func(i interface{}) {

		switch i.(type) {
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("boolean")
		case int:
			fmt.Println("integer")
		default:
			fmt.Println("other")
		}

	}

	// value("madhavaas")
	value(true)
}
