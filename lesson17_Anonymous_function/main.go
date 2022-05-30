package main

import "fmt"

//関数
//無名関数

//簡易的に関数宣言したいときに無名関数が用いられることがある

func main() {
	f := func(x, y int) int {
		return x + y
	}

	i := f(1, 2)
	fmt.Println(i)

	i2 := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(i2)

}
