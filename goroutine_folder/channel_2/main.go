package main

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum //チャネルに送信する
	}
	// c <- sum //チャネルに送信する
	close(c)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, 2)
	go goroutine1(s, c)
	// goroutineでループを回した時にclose()しないとdeadlock
	// 随時main関数にgoroutine1内の処理結果を返すことができる
	for i := range c {
		fmt.Println(i)
	}
}
