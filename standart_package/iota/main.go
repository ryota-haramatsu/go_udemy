package main

import "fmt"

const (
	// c1 = iota // 自動連番
	// c2 = iota
	// c3 = iota

	// こちらでもOK
	c1 = iota // 自動連番
	c2
	c3
)

const (
	// KB, MB, GBの時に使用可能
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

func main() {
	fmt.Println(c1, c2, c3) // 0 1 2と出力
	fmt.Println(KB, MB, GB)
}
