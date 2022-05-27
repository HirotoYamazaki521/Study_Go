package main

import "fmt"

//型
//文字列型

func main() {
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	//改行したい場合
	fmt.Println(`test
	test
		test`)

	//""を文字列とし扱う
	fmt.Println("\"")
	fmt.Println(`"`)

	//文字列はバイト配列の集まり
	fmt.Println(string(s[0]))
}
