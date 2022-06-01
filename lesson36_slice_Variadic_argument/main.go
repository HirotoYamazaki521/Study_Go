package main

import "fmt"

//スライス
//可変長引数

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	fmt.Println(Sum())

	sl := []int{1, 2, 3} //sliceとなって1,2,3が渡される
	fmt.Println((Sum(sl...)))
	//引数の数を指定しないで関数に渡すことができる
}
