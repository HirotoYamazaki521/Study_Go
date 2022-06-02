package main

import "fmt"

//独自型

//MyIntという型をint型で独自の型として宣言できる
type MyInt int

//独自型にメソッドを作ることもできる
func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)

	//i := 100
	//MyInt型とint型では計算ができない
	//fmt.Println(mi + i)
	//あるデータの型を厳しく設定することで、他のデータ型と計算できないようにできる

	mi.Print()

}
