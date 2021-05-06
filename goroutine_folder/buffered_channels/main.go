package main

import "fmt"

func main() {
	ch := make(chan int, 2) // buffered channels
	ch <- 100
	// fmt.Println(len(ch))
	ch <- 200
	// fmt.Println(len(ch))
	// x := <-ch
	// fmt.Println(x)
	// fmt.Println(len(ch))

	close(ch) // ループを回す際にclose
	for c := range ch {
		fmt.Println(c)
	}
}
