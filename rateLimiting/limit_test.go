package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"testing"
	"time"
)

var (
	requestChan  = make(chan []int)
	responseChan = make(chan string)
	wg           sync.WaitGroup
)

func sendRequest(t *testing.T) {
	for payload := range requestChan {
		for x := 1; x <= payload[0]; x++ {
			go func() {
				// defer wg.Done()
				resp, err := http.Get("http://127.0.0.1:8080/ping")
				if err != nil {
					t.Error(err)
				}
				//We Read the response body on the line below.
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					t.Error(err)
				}
				sb := "Goroutine " + strconv.Itoa(payload[1]) + " : " + string(body)
				responseChan <- sb
			}()
		}
	}
}

func TestServer(t *testing.T) {

	gap := 500
	ticker := time.NewTicker(time.Duration(gap) * time.Millisecond)

	go sendRequest(t)

	go func() {
		for r := range responseChan {
			fmt.Println(r)
		}
	}()

	counter := 0
	for {
		<-ticker.C
		counter += 1
		requestChan <- []int{2, counter}
		// wg.Add(1)
	}

	// go func() {
	// 	wg.Wait()
	// 	close(responseChan)
	// }()

}
