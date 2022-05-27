package main //packageの宣言は１つだけ

import (
	"fmt"
) //format packageでprint println等のパッケージ
//Helloworld
/*
複数行の
コメントアウト
*/

//i5 := 500　暗黙的な定義は関数外でしていできない
var i5 int = 500

//基本的には型指定を行う明示的な定義を使う

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	//明示的な定義
	//var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)

	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3) //型の初期値が入る int = 0 ,文字列 = 空文字

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)
	//fmt.Println("Hello World")
	//fmt.Println(time.Now())

	i = 150
	fmt.Println(i)

	//暗黙的な定義
	//変数名 :=値　型指定が必要ない
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	//i4 := 500
	//fmt.Println(i4)

	//i4="Hello"
	//fmt.Println(i4)  型推論といって最初に定義した値で型を決める

	fmt.Println(i5)

	outer()

	/*
		fmt.Println(s4)
		main関数の中にs4がないので、
		outer関数の中でのみ有効
	*/

	/*
		goは宣言した変数を必ず使う必要がある
	*/
	var s5 string = "Not Use"

	fmt.Println(s5)
}
