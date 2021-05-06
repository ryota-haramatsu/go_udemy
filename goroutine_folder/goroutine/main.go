package main

import (
	"fmt"
	"sync"
)

// - goroutine (go)
// - sync.WaitGroup wg.Add() wg.Wait() wg.Done()

func goroutine(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go goroutine("world", &wg) // goと追加するだけで並列処理できる
	normal("hello")
	wg.Wait()
}
