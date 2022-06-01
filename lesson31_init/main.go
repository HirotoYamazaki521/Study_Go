package main

import "fmt"

//init
//パッケージの初期化を目的とした特殊な関数

//init関数はmain関数より早く実行される
//mainルーチンより前に初期化処理を確実に実行する
func init() {
	fmt.Println("init")
}

func init() {
	fmt.Println("init2") //init関数は複数にわけるよりはまとめたほうがよい
}

func main() {
	fmt.Println("Main")
}
