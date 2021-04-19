package main

import "fmt"

func main() {
	var c = make([]int, 0, 5)
	var c = make([]int, 5) // 0埋めされたcap5, len5のスライス
	fmt.Println(c)
	for i := 0; i < 5; i++{
		c = append(c, i)
		// fmt.Println(c)
	}
	fmt.Println(c)
}
