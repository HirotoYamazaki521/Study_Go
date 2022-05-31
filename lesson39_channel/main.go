package main

import "fmt"

//channel
//複数のゴルーチン間でのデータの受け渡しをするために設計されたデータ構造。
//宣言、操作

func main() {
	//チャネルを指定しない場合は双方向としてのチャネルとして機能
	var ch1 chan int //int型のデータを保持できるチャネルの作成

	//受信専用のチャネルを明示的に宣言
	//var ch2 <-chan int

	//送信専用
	//var ch3 chan<- int

	//make関数を使うことでチャネルとしての機能
	//書き込み、読み込み機能をもたせる
	ch1 = make(chan int)

	ch2 := make(chan int)

	//データの容量(バッファサイズ)を調べる
	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	//バッファサイズを指定して作成
	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	//チャネルからデータを受信したり、チャネルのデータを更新する
	//データをチャネルに送信
	ch3 <- 1
	//データの要素数を調べる
	fmt.Println(len(ch3))

	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3))

	//チャネルからデータを受信
	i := <-ch3
	fmt.Println(i)
	fmt.Println("len", len(ch3))

	i2 := <-ch3
	fmt.Println(i2)
	fmt.Println("len", len(ch3))

	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3)) //チャネルはキューの性質を持っている先入れ先だし

	//バッファサイズを超えたデータの取り出し
	ch3 <- 1
	fmt.Println("len", len(ch3))
	fmt.Println(<-ch3) //途中で受信するとデッドロックがなくなる
	fmt.Println("len", len(ch3))
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6
	//バッファサイズを超えたデータを送信するとデッドロックが起きる

}
