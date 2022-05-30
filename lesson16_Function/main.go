package main

import "fmt"

//関数

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
	//引数と返り値を見ることでなにをする関数かが分かりやすい
}

//返り値のない関数
func Noreturn() {
	fmt.Println("No Return")
	return
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(9, 4)
	fmt.Println(i2, i3)
	//変数宣言されたものは必ず使わないとエラーが起こる
	//fmt.Println(i2)

	//使いたくない場合
	//i2,_:=Div(9,4)

	i4 := Double(1000)
	fmt.Println(i4)

	Noreturn()
}
