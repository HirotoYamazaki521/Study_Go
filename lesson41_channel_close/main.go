package main

import (
	"fmt"
	"time"
)

//channel
//close

//チャネルはクローズという状態を持っている
//組み込み関数makeで生成したチャネルはオープン状態から始まるが
//送受信を終えたチャネルを明示的に閉じることも可能

func reciever(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	ch1 := make(chan int, 2)

	/*
		ch1 <- 1

		close(ch1)

		//閉じたチャネルにはデータの送信ができない
		//ch1 <- 1

		//データの受信は可能
		//fmt.Println(<-ch1)

		i, ok := <-ch1
		//第2変数にはチャネルのオープン状態を表す真理値が出力
		//チャネルのバッファ内が空でcloseの場合に、false
		fmt.Println(i, ok)

		i2, ok := <-ch1
		fmt.Println(i2, ok)
	*/

	//どのレシーバーが動くかは、実行結果によって変わる
	go reciever("1.goroutin", ch1)
	go reciever("2.goroutin", ch1)
	go reciever("3.goroutin", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
	//チャネルが閉じてからゴルーチンの完了を待ちたい(簡易的)
	time.Sleep(3 * time.Second)
}
