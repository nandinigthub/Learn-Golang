package advance

import (
	"fmt"
	"net/url"
)

const myurl = "https://datausa.io/api/data?drilldowns=Nation&measures=Population"

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Urrls() {
	// response, err := http.Get(myurl)
	// checkNilErr(err)

	// fmt.Printf("response type: %T \n", response)

	// datatype, err := ioutil.ReadAll(response.Body)
	// checkNilErr(err)

	// content := string(datatype)
	// fmt.Println(content)

	// parsing
	res, err := url.Parse(myurl)
	checkNilErr(err)

	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	// fmt.Println(res.RawQuery)
	// fmt.Println(res.Port())

	// qparams := res.Query()
	// // it is of form key value pair = url.Values
	// fmt.Printf("query or parameter type : %T \n", qparams)

	// for _, val := range qparams {
	// 	fmt.Println("param vlaue", val)
	// }

	// create url - https://dog.ceo/api/breeds/image/random

	newurl := &url.URL{
		Scheme: "https",
		Host:   "dog.ceo",
		Path:   "api/breeds/image/random/",
	}

	createurl := newurl.String()
	fmt.Println("my created url:\n", createurl)

}
