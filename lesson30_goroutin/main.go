package main

import (
	"fmt"
	"time"
)

//go goroutin
//goの並行処理の機能
//goroutin=スレッドよりも小さい処理単位
//go文を使うことで暗黙的に並行処理を行う

//サブとして並行で走らせる関数
func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
		//一定時間感覚を開けて出力
	}
}

func main() {
	go sub()
	go sub()

	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}
}
