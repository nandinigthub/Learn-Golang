package basics

import "fmt"

// enumerated types - basically helps in managing options in group
// not a type in golang but we create it using type and const

type orderstatus string

const (

	// iota for int, string for string
	// Received orderstatus = iota
	Received orderstatus = "recieved"
	updated  orderstatus = "updated"
	// updated
	// confirmed
	// delivered
	// refund
)

func changeorder(status orderstatus) {
	fmt.Println("change order status", status)
}
func Enum() {
	changeorder(updated)
}
