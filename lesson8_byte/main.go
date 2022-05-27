package main

import "fmt"

//型
//byte(uint8)型

func main() {
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	//文字列はアスキーコードで表現されているため
	fmt.Println(string(byteA))

	//バイトのスライスを文字列に変換
	c := []byte("HI")
	fmt.Println(c)

	//文字列に変換
	fmt.Println(string(c))
}
