// hitesh choudhary vdo
package basics

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex
var signals = []string{"test"}

func Mutexs() {
	// mutex lock like os same read write lock unlock thing

	fmt.Println("executing mutex")

	websites := []string{
		"https://go.dev",
		"https://google.com",
		"https://github.com",
		"https://fb.com",
		"https://shaadi.com",
	}

	for _, web := range websites {
		go getStatuscode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

func getStatuscode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("errrrrrrrrr for", endpoint)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d stautus code is for %s \n", res.StatusCode, endpoint)
	}

}
