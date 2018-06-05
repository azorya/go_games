package main

import (
	"fmt"
	"time"
)

func main() {
	var ping = make(chan struct{})
	var pong = make(chan struct{})
	var done = make(chan struct{})

	type result struct {
		count int64
		id    bool
	}

	r := make(chan result, 2)

	go func() {
		time.Sleep(1 * time.Second)
		close(done)
	}()

	go func() {
		var iping int64

		for {

			iping++
			select {
			case ping <- struct{}{}:
				<-pong
			case <-done:
				close(ping)
				r <- result{iping, true}
				return
			}
		}
	}()

	go func() {
		var ipong int64

		for {
			ipong++
			select {
			case <-ping:
				pong <- struct{}{}
			case <-done:
				close(pong)
				r <- result{ipong, false}
				return
			}
		}
	}()

	v1 := <-r
	v2 := <-r
	close(r)
	fmt.Println(v1, v2)
}
