package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("表示")

	/*
		//fmt標準
		print("Hello\n")	//\nで改行
		//改行
		println("Hello!")
	*/

	/*
		//Println系 各々の文字列は半角スペースで区切られ、文字列の最後に改行
		fmt.Println("Hello", "world!")
		fmt.Println("Hello", "world!")
	*/
	/*
		//Printf系 フォーマットを指定
		fmt.Printf("%s\n", "Hello")	//%s 文字列の書式指定
		fmt.Printf("%#v\n", "Hello")//%#v 値の文法をそのまま出力
	*/

	/*
		//Sprint系 出力ではなくフォーマットした結果を文字列で返す。
		s := fmt.Sprintf("Hello")
		s1 := fmt.Sprintf("%v\n", "Hello")
		s2 := fmt.Sprintln("Hello")

		fmt.Println(s)
		fmt.Println(s1)
		fmt.Println(s2)
	*/
	//Fprint系　書き込み先指定 任意のio writer型への出力関数
	fmt.Fprint(os.Stdout, "Hello") //os.Stdout 標準出力
	fmt.Fprintf(os.Stdout, "Hello")
	fmt.Fprintln(os.Stdout, "Hello")

	f, _ := os.Create("test.txt")
	defer f.Close()

	fmt.Fprintln(f, "Fprint")
}
