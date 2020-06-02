package main

import (
		"net/http"
		"fmt"
		//"io/ioutil"
		"sync"
		"golang.org/x/net/http2"
)

//body, err := ioutil.ReadAll(resp.Body)
//if err != nil {
//	// handle error
//}
//fmt.Println(string(body))

type Work struct {
	method string
	data string
	n int
}

func (w *Work) runWorker(client *http.Client) {

	req, err := http.NewRequest(w.method, "http://example.com", nil)
	if err != nil {
		fmt.Println("did not create new request")
	}
	for i := 0; i < w.n; i++ {
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("did not do request")
		}
		fmt.Println(resp.Status, "did ", w.method)
	}
}

func makeConcurrentRequests() {
	var wg sync.WaitGroup
	wg.Add(4) // sum of concurrent workers

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         b.Request.Host,
		},
		MaxIdleConnsPerHost: min(b.C, maxIdleConn),
		DisableCompression:  b.DisableCompression,
		DisableKeepAlives:   b.DisableKeepAlives,
		Proxy:               http.ProxyURL(b.ProxyAddr),
	}


	client := &http.Client{}
	wGET := &Work{
		method: "GET",
		data: "",
		n : 2,
	}
	wPOST := &Work{
		method: "POST",
		data: "",
		n : 3,
	}

	for i := 0; i < 2; i++ { // concurent workers
		go func() {
			wGET.runWorker(client)
			wg.Done()
		}()
	}

	for i := 0; i < 2; i++ { // concurent workers
		go func() {
			wPOST.runWorker(client)
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {

	makeConcurrentRequests()

}
