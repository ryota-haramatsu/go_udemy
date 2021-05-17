package main

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //チャネルに送信する
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //チャネルに送信する
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, 2)
	go goroutine1(s, c)
	go goroutine2(s, c)
	x := <-c // チャネル受信 sumが入るまでwait
	fmt.Println(x)
	y := <-c
	fmt.Println(y)
}
