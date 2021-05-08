package main

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("スタート")
	time.Sleep(2 * time.Second)
	fmt.Println("フィニッシュ")
	ch <- "結果"
}

func main() {
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	go longProcess(ctx, ch)
CTXLOOP:
	for {
		select {
		// 設定時間経ってもgoroutineが終了しなければ終了
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <-ch:
			fmt.Println("成功")
			break CTXLOOP
		}
	}
	fmt.Println("###############")
}
