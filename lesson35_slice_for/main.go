package main

import "fmt"

//スライス
//for

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	/*
		for i, v := range sl {
			fmt.Println(i, v) //インデックス番号と要素を順番に表示
		}

		for _, v := range sl { //インデックス番号を使わずに表示、逆も可
			fmt.Println(v)
		}

		for i := range sl { //インデックス番号を使わずに表示、逆も可
			fmt.Println(i, sl[i])
		}
	*/

	for i := 0; i < len(sl); i++ { //インデックス番号にあった要素を出力できる
		fmt.Println(sl[i])
	}

}
