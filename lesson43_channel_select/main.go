package main

import "fmt"

//channel
//select

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	/*
		ch2 <- "A"
		ch1 <- 1
		ch2 <- "B"
		ch1 <- 2
	*/

	/*
		v1 := <-ch1	//空のチャネルから値を受信できずデッドロック
		v2 := <-ch2	//ここまで到達できない

		fmt.Println(v1)
		fmt.Println(v2)
	*/

	//チャネルの1と2に入った場合の処理の分岐処理
	//switch文と違い、最初に成立したケースが優先されるのではなく、
	//どのケース式が実行されるかはランダム

	/*
		参考：https://code-graffiti.com/channnel-and-select-in-golang/

		ケースの１つが実行可能になるまで他のケースをブロックし、
		その後に他ケースを実行します。
		複数のケースが準備できている場合は、ランダムに1つを選択して順に処理。
	*/

	select {
	case v1 := <-ch1: //チャネルに対する処理以外はエラーになる
		fmt.Println(v1 + 1000)
	case v2 := <-ch2:
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("どちらでもない")
	}

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	//reciever
	go func() {
		for {
			i := <-ch3
			ch4 <- i * 2
		}
	}()

	go func() {
		for {
			i2 := <-ch4
			ch5 <- i2 - 1
		}
	}()

	//各々、非同期に処理されるチャネルのデータを適切に処理できる
	n := 0
L:
	for {
		select {
		case ch3 <- n:
			n++
		case i3 := <-ch5:
			fmt.Println("received", i3)
		default:
			if n > 100 {
				break L
			}
		}
		/*
			if n > 100 {
				break
			}
		*/
	}

}
