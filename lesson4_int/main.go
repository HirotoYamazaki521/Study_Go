package main

import "fmt"

//型
//整数型

func main() {
	var i int = 100

	/*
		最大値 最小値
			int8	8-bit integers (-128 to 127)
			int16	8-bit integers (-32768 to 32767)
			int32	8-bit integers (-2147483648 to -2147483647)
			int64	8-bit integers (-9223372036854775808 to -9223372036854775807)

		PCの環境に依存する
	*/

	var i2 int64 = 200

	fmt.Println(i + 50)

	//fmt.Println(i + i2) 同じint型でも計算できない
	//明示的に宣言したint型と完共依存のint型ではまったく別の型と認識される

	fmt.Printf("%T\n", i2) //%Tは型を表示する書式

	fmt.Printf("%T\n", int32(i2)) //型変換

	fmt.Println(i + int(i2))
}
