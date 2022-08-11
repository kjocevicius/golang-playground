package main

import (
    "fmt"
)

func fSimple(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func fc(from string, channel chan string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
    channel <- "Done!"
}

func main() {
    channel := make(chan string)

    fSimple("direct")

    go fc("goroutine", channel)

    go func(msg string) {
        fmt.Println(msg)
    }("going")

    msg := <-channel // No need to sleep - 'awaits' for value
    fmt.Println(msg)
}