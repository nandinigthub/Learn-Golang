package basics

import (
	"fmt"
	"strconv"
)

func value(i interface{}) {
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

func Help() {
	var input string
	fmt.Println("Enter a value:")
	fmt.Scan(&input)

	// Try to infer the type
	if val, err := strconv.Atoi(input); err == nil {
		value(val) // Input is an integer
	} else if val, err := strconv.ParseBool(input); err == nil {
		value(val) // Input is a boolean
	} else {
		value(input) // Default to string
	}
}
