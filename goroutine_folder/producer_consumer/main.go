package main

import (
	"fmt"
	"sync"
	"time"
)

// producerとconsumer

func producer(ch chan int, i int) {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		// fmt.Println("process", i*1000)
		// wg.Done() //一つ処理が終わるごとにDone
		func() {
			fmt.Println("process", i*1000)
			wg.Done()
		}()
	}
	fmt.Println("###########")
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Producer
	// ループで10回producerを作成して、10回Done()する必要
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	// Consumer
	go consumer(ch, &wg)
	wg.Wait()
	close(ch)
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
