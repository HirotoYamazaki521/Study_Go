package main

import "fmt"

//関数
//関数を引数にとる関数

func CallFunction(f func()) { //引数で取った関数を中で実行
	f()
}

func main() {
	CallFunction(func() {
		fmt.Println("I'm function")
	})
}
