package main

import (
	"fmt"
	"sync"
	"time"
)

// https://gobyexample.com/waitgroups
// https://gobyexample.com/channels

const SIZE = 5
const LOOPS_IN_GOROUTINE = 3

func f(from string, wg *sync.WaitGroup, resultChannel chan string, modifier int) {
	defer wg.Done()
	for i := 0; i < LOOPS_IN_GOROUTINE; i++ {
		time.Sleep((modifier * i ) * time.Second)
		result := fmt.Sprintf("%s -> %d", from, i)
		resultChannel <- result
	}
}

func main() {

	var wg sync.WaitGroup
	channelSize := SIZE * LOOPS_IN_GOROUTINE
	resultChannel := make(chan string, channelSize)

	for i := 0; i < SIZE; i++ {
		wg.Add(1)
		go f(fmt.Sprintf("goroutine %d", i), &wg, resultChannel, i*100)
	}

	wg.Wait()
	close(resultChannel)
	fmt.Println("done")

	for result := range resultChannel {
		fmt.Println(result)
	}
}
