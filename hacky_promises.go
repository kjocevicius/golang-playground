package main

import (
	"fmt"
	"time"
)

type PromisedFunction[T any] func() T

type Promise[T any] struct {
	ch chan T
}

func AsPromise[T any](fn PromisedFunction[T]) *Promise[T] {
	result := new(Promise[T])
	result.ch = make(chan T, 1)
	go func(resultChan chan T, f PromisedFunction[T]) T {
		out := f()
		resultChan <- out
		return out
	}(result.ch, fn)

	return result
}

func (promise Promise[T]) Await() T {
	result := <-promise.ch
	return result
}

func AwaitAll[T any](promises []Promise[T]) []T {
	results := make([]T, len(promises))
	for idx, promise := range promises {
		results[idx] = promise.Await()
	}
	return results
}

func main() {
	trySinglePromise()
	tryMultiplePromises()
}

func trySinglePromise() {
	promise := somePromise(0)
	result := promise.Await()
	fmt.Println(result)
}

func tryMultiplePromises() {
	count := 100
	promises := make([]Promise[int], count)
	for i := 0; i < count; i++ {
		promise := somePromise(i)
		promises[i] = *promise
	}
	result := AwaitAll(promises)
	fmt.Println(result)
}

func somePromise(modifier int) *Promise[int] {
	a := 1000000
	b := 22222
	promise := AsPromise(func() int {
		result := a + b + modifier
		time.Sleep(2 * time.Second)
		return result
	})
	return promise
}
