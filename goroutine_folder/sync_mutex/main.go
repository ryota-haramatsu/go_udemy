package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

// 一方のgoroutineが書き込みを行っている時にもう一方が書き込みできないようにする
func (c *Counter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	// c := make(map[string]int)
	c := Counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 10; i++ {
			// c["key"] += i
			c.Inc("Key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			// c["key"] += i
			c.Inc("Key")
		}
	}()
	time.Sleep(1 * time.Second)
	// fmt.Println(c, c["key"])
	fmt.Println(c, c.Value("Key"))
	// fmt.Println(runtime.NumGoroutine())
}
