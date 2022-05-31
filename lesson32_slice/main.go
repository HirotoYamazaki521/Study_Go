package main

import "fmt"

//スライス
//宣言、操作

func main() {
	var sl []int //[]に要素数を宣言しない
	fmt.Println(sl)

	//明示的宣言
	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	//暗黙的宣言
	sl3 := []string{"A", "B"}
	fmt.Println(sl3)

	sl4 := make([]int, 5) //第2引数に要素数
	fmt.Println(sl4)

	//値の挿入
	sl2[0] = 1000
	fmt.Println(sl2)

	//値の取り出し
	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)

	fmt.Println(sl5[0])

	fmt.Println(sl5[2:4]) //2番から4番の手前まで取得

	fmt.Println(sl5[:2]) //2番の手前まで

	fmt.Println(sl5[2:]) //2番から最後まで

	fmt.Println(sl5[len(sl5)-1]) //最後の要素の取り出し

	fmt.Println(sl5[1 : len(sl5)-1]) //最初と最後をのぞいたすべての要素

}
