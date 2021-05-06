package main

import (
	"fmt"
	"time"
)

// select{}

// お互いがブロッキングしないように
func goroutine1(ch chan string) {
	// 無限にパケットを渡す想定
	for {
		ch <- "1からのパケット"
		time.Sleep(1 * time.Second)
	}
}

func goroutine2(ch chan string) {
	// 無限にパケットを渡す想定
	for {
		ch <- "2からのパケット"
		time.Sleep(2 * time.Second)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go goroutine1(c1)
	go goroutine2(c2)
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
