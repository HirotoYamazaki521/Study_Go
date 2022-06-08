package main

import (
	"context"
	"fmt"
	"time"
)

//cotnext
//APIのサーバーやクライアントを使うときに
//contextを提供してキャンセルやタイムアウトをできる仕組み

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("開始")
	time.Sleep(2 * time.Second)
	fmt.Println("終了")
	ch <- "実行結果"
}

func main() {
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second) //ctxに1秒間を設定
	//ctx, cancel := context.WithTimeout(ctx, 3*time.Second) //成功する
	defer cancel() //リソースの解放処理

	go longProcess(ctx, ch) //

	//cancel()

L:
	for {
		select {
		case <-ctx.Done():
			//ctxで1秒間で処理が終わらなかったら動作
			fmt.Println("##########Error###########")
			fmt.Println(ctx.Err())
			break L
		case s := <-ch:
			fmt.Println(s)
			fmt.Println("success")
			break L
		}

	}

	fmt.Println("ループ抜けた")
}
