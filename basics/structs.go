package basics

import (
	"fmt"
	"time"
)

// struct order
type customer struct {
	fname string
	lname string
}
type order struct {
	id     int
	name   string
	status bool
	time   time.Time
	customer
}

// create constructor
// struct k neeche hi banate hai - sath m chipka k
// naming convention hai starts with New (n capital m)

// func NewOrder(id int, name string, status bool) *order {
// 	// setup struct
// 	myOrders1 := order{
// 		id:     id,
// 		name:   name,
// 		status: status,
// 	}

// 	return &myOrders1

// }

// func using struct -> receivers func
// to update use pointer
// by convention we use struct first letter as alias shows its a receiver func
// func (o *order) changename(name string) {
// 	o.name = name
// }

// since we are not modifying the input no need to use pointer
// func (o order) getstatus() bool {
// 	return o.status
// }

func Stucs() {
	// myOrders1 := order{
	// 	id:     1,
	// 	name:   "xyz",
	// 	status: false,
	// }

	// myOrders1.time = time.Now()
	// // get field
	// // fmt.Println(myOrders1.status)
	// fmt.Println("order 1", myOrders1)

	// myOrders2 := order{
	// 	id:     2,
	// 	name:   "abc",
	// 	status: true,
	// 	time:   time.Now(),
	// }
	// fmt.Println("order 2", myOrders2)

	// myOrders1.changename("riddhi")
	// fmt.Println(myOrders1.getstatus())
	// fmt.Println("order1 updated", myOrders1)

	// directly declare struct for one time use

	// lang := struct {
	// 	name string
	// 	know bool
	// }{"golang", true}
	// fmt.Println("languages", lang)

	// // constructor func call
	// myOrder := NewOrder(1, "charu", true)
	// fmt.Println(myOrder)
	// // fmt.Println(myOrder.name)

	// myOrder2 := NewOrder(2, "mehek", false)
	// fmt.Println(myOrder2)

	// --------------------------------------------------->

	// stuct embedding
	// m1

	mycustomer := customer{
		fname: "Jaya",
		lname: "balani",
	}

	myOrders1 := order{
		id:       1,
		name:     "xyz",
		status:   false,
		time:     time.Now(),
		customer: mycustomer,
	}

	fmt.Println(myOrders1)

	//m2
	myOrders2 := order{
		id:     2,
		name:   "abc",
		status: false,
		customer: customer{
			fname: "rekha",
			lname: "godani",
		},
	}

	fmt.Println(myOrders2)

}
