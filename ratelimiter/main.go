package main

import (
	"fmt"
	"time"
)

func main() {
	request := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		request <- i
	}
	close(request)

	limiter := time.Tick(1000 * time.Millisecond)
	for req := range request {
		<-limiter
		fmt.Println("Request: ", req, time.Now())
	}

	rafaga := make(chan time.Time, 3)
	go func() {
		for t := range time.Tick(1000 * time.Millisecond) {
			for i := 0; i < 3; i++ {
				rafaga <- t
			}
		}
	}()

	rafagaRequest := make(chan int, 15)
	for i := 1; i <= 15; i++ {
		rafagaRequest <- i
	}
	close(rafagaRequest)
	for req := range rafagaRequest {
		<-rafaga
		fmt.Println("Request: ", req, time.Now())
	}
}
