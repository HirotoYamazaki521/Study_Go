package main

import (
	"fmt"
)

//定数
//頭文字を大文字にすると他のパッケージからも呼び出せる
const Pi = 3.14

const (
	URL      = "http://xxx.co.jp"
	SiteName = "test"
)

//値を省略して定義することも可能
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

//定数の場合は、型の最大値を超えて定義可能
//var Big int = 9223372036854775807 + 1
const Big = 9223372036854775807 + 1

const (
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(Pi)

	//Pi = 3
	//fmt.Println(Pi)
	fmt.Println(Pi, SiteName)
	fmt.Println(A, B, C, D, E, F)
	fmt.Println(Big - 1)
	fmt.Println(c0, c1, c2)
}
