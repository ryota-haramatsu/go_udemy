package main

import (
	"fmt"
	"gomodtest/mylib"
)

func main() {
	s := []int{1, 2, 2, 45, 100}
	mylib.Average(s)
	fmt.Println(mylib.Average(s))
}
