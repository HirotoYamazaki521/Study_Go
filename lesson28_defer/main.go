package main

import (
	"fmt"
	"os"
)

//defer
//関数の終了時に実行される処理を登録することができる

func TestDefer() {
	defer fmt.Println("END") //最後に実行される
	fmt.Println("START")
}

//同じdeferでも後から登録された式から順に評価する
func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func main() {
	TestDefer()

	/*
		defer func() { //無名関数
			fmt.Println("1")
			fmt.Println("2")
			fmt.Println("3")
		}()
	*/

	RunDefer()

	//defer文を使ったリソースの解放処理
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() //開いたファイルを閉じる

	file.Write([]byte("Hello")) //ファイルに文字列を書き込む

}
