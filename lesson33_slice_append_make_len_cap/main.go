package main

import "fmt"

//スライス
//append make len cap

func main() {
	sl := []int{100, 200}
	fmt.Println(sl)

	//append
	sl = append(sl, 300) //追加したいスライス,追加したい要素
	fmt.Println(sl)

	sl = append(sl, 400, 500, 600)
	fmt.Println(sl)

	//make スライスを作成する関数
	sl2 := make([]int, 5)
	fmt.Println(sl2)

	//len
	fmt.Println(len(sl2))

	//キャパシティを調べる キャパシティ=容量
	fmt.Println(cap(sl2))

	sl3 := make([]int, 5, 10) //make関数、第2引数が要素数、第3引数が容量
	//プログラムのメモリを気にする場合

	fmt.Println(len(sl3))

	fmt.Println(cap(sl3))

	sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)

	fmt.Println(len(sl3))

	fmt.Println(cap(sl3))

	//容量を超えた要素数になると容量が倍になってしまう

}
