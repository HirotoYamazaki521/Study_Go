package main

import "fmt"

//if
//条件関数

func main() {
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don't know")
	}

	//簡易条件文　条件式の前に簡易文を書ける
	if b := 100; b == 100 {
		fmt.Println("one handred")
	}

	x := 0
	if x := 2; true {
		fmt.Println(x)
	}
	fmt.Println(x)
	//if文内部の変数が優先される
	//if文の中の変数は外側からアクセスできない
}
