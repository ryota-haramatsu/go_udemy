package main

import "fmt"


func foo(params ...int){
	fmt.Println(len(params), params)
	for _, param := range params{
		fmt.Println(param)
	}
}

func main() {
	foo(1,2,3,4,5)

	s := []int{1,2,3} // 
	fmt.Println(s)
	foo(s...)
}
