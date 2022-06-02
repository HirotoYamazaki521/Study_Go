package main

import (
	"fmt"
	"time"
)

//channel
//チャネルとゴルーチン

func reciever(c chan int) {
	for { //チャネルに値が渡されたらデータを出力する
		i := <-c
		fmt.Println(i)
	}
}

func main() {

	/*
		ch1 := make(chan int)
		fmt.Println(<-ch1)
		チャネルはゴルーチン間でデータを共有するための仕組みなので、
		必然的に複数のゴルーチン間でチャネルを共有するというプログラムが必要
	*/

	ch1 := make(chan int)
	ch2 := make(chan int)

	go reciever(ch1)
	go reciever(ch2)

	i := 0
	for i < 100 {
		ch1 <- i
		ch2 <- i
		time.Sleep(50 * time.Millisecond)
		i++
	}
}
