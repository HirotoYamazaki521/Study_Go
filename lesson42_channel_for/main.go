package main

import "fmt"

//channel
//for

//簡易式のfor文を使うことができる
func main() {
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)
	for i := range ch1 {
		fmt.Println(i)
		//3番まで値を受け取ったら次の値がないので、
		//空のチャネルから値を受け取り、デッドロックが発生する
		//そのため、チャネルでfor文を使う場合はcloseする必要あり

	}
}
