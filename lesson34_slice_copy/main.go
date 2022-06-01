package main

import "fmt"

//slice
//copy

func main() {
	/*
		sl := []int{100, 200}
		sl2 := sl

		sl2[0] = 1000
		fmt.Println(sl)
		fmt.Println(sl2)
		//slの指す0番とsl2の指す0番は同じなので、両者の0番値が更新されてしまう参照型特有の性質

		var i int = 10
		i2 := i
		i2 = 100
		fmt.Println(i, i2)
		//基本型は別物なので、同時に更新されない
	*/

	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)
	fmt.Println(sl2)
	n := copy(sl2, sl) //第1引数にコピー先、第2引数にコピー元 nはコピーに成功した数

	fmt.Println(n, sl2)
}
