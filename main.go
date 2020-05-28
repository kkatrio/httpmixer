package main

import (
		"net/http"
		"fmt"
		//"io/ioutil"
		"sync"
)

//body, err := ioutil.ReadAll(resp.Body)
//if err != nil {
//	// handle error
//}
//fmt.Println(string(body))


func runWorker(client *http.Client, nReqs int, method string) {

	req, err := http.NewRequest(method, "http://example.com", nil)
	if err != nil {
		fmt.Println("did not create new request")
	}
	for i := 0; i < nReqs; i++ {
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("did not do request")
		}
		fmt.Println(resp.Status, "did ", method)
	}
}

func makeConcurrentRequests() {
	var wg sync.WaitGroup
	wg.Add(4) // sum of concurrent workers

	client := &http.Client{}

	nGETReqs := 10
	for i := 0; i < 2; i++ { // concurent workers
		go func() {
			runWorker(client, nGETReqs, "GET")
			wg.Done()
		}()
	}

	nPOSTReqs := 5
	for i := 0; i < 2; i++ { // concurent workers
		go func() {
			runWorker(client, nPOSTReqs, "POST")
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {

	makeConcurrentRequests()

}
