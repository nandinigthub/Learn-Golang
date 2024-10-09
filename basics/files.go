package basics

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Files() {
	fmt.Println("file handling in go.....")

	content := "this is the content of file"

	file, err := os.Create("./myfile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("length is :", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)
	checkNilErr(err)

	fmt.Print("text data in file: \n", string(databyte))
}
