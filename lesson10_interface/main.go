package main

import "fmt"

//interface & nil
//あらゆる型と互換性のある特殊な型

//nil = 値を何も持っていない

func main() {
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)

	x = 3.14
	fmt.Println(x)

	x = "A"
	fmt.Println(x)

	x = [3]int{1, 2, 3}
	fmt.Println(x)

	//データ特有の計算や、演算ができない
	x = 2
	fmt.Println(x + 3)
}
