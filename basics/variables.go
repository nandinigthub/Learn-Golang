package basics

import "fmt"

func Vari() {
	var username string = "nandini"
	fmt.Println(username)
	fmt.Printf("variable type: %T \n", username)

	var age int = 20
	fmt.Println(age)

	var name = "golang"
	fmt.Println(name)

	// usually written in uppercase letters
	const (
		PORT  = 400
		CLASS = "xii"
		LNAME = "go"
	)

	// multi variable declaration
	var a, b = 6, "Hello"
	c, d := 7, "World!"
	// single line print
	fmt.Println(a, b, c, d)

	fmt.Println(PORT)
	fmt.Println(CLASS)
	fmt.Println(LNAME)
}
