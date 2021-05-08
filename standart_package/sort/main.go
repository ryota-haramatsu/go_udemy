package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{1, 3, 4, 5}
	s := []string{"apple", "orange", "banana"}
	p := []struct {
		Name string
		Age  int
	}{
		{"A", 30},
		{"D", 20},
		{"B", 10},
		{"C", 50},
	}
	// 整数ソート
	sort.Ints(i) 
	fmt.Println(i)

	// 文字列ソート
	sort.Strings(s) 
	fmt.Println(s)

	// 構造体ソート
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	fmt.Println(p)
}
