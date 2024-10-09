package basics

import (
	"fmt"
	"maps"
)

func Mapii() {
	// key value
	m := make(map[string]string)

	m["name"] = "golang"
	m["age"] = "20"
	m["class"] = "xii"
	// fmt.Println(m["name"], m["age"])
	// fmt.Println(len(m))
	// fmt.Println(m)

	// // The make()function is the right way to create an empty map.
	// //  If you make an empty map in a different way and write to it, it will causes a runtime panic.
	// m2 := make(map[string]int)
	// m2["area"] = 3000
	// m2["dept"] = 2
	// m2["price"] = 400
	// m2["numver"] = 3456789
	// fmt.Println(m2)
	// fmt.Println(m2["phone"])
	// delete(m2, "price")
	// fmt.Println(m2)
	// fmt.Println(len(m2))

	// //directly intialise
	// m3 := map[string]int{"price": 40, "age": 30}
	// fmt.Println(m3)

	// to print value and check key present
	val, key := m["age"]
	// only checking key exists or not
	_, key2 := m["class"]
	fmt.Println(val)
	fmt.Println(key)
	fmt.Println(key2)
	fmt.Println("maps ex")

	// just like slice to check whether equal or not
	m3 := make(map[string]string)
	m3["area"] = "3000"
	m3["dept"] = "3"
	m3["price"] = "400"
	m3["numver"] = "3456789"

	fmt.Println(maps.Equal(m, m3))

}
