package main

import "fmt"

//switch
//型スイッチ
//型アサーション 動的に変数の型をチェっクする

func anything(a interface{}) {
	fmt.Println(a)
}

func main() {
	anything("aaa")
	anything(1)
}
