package main

import "fmt"

//Generics Go1.18から追加された機能
//型を抽象化してコードの再利用性を向上させつつ、静的型付け言語の乙、型安全性を高める

/*
func PrintSliceInts(i []int) {
	for _, v := range i {
		fmt.Println(v)
	}
}

//通常はPrintSliceInts関数にintのスライス以外の引数を渡す場合はもう１つ関数を作る必要がある
func PrintSliceStrings(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}
*/

//Genericを使って抽象化
//any型とは、従来のinterface型のエイリアスとなっておりあらゆる型を入れることができる
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	//PrintSliceInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	//PrintSliceStrings([]string{"a", "b", "c", "d"})
	PrintSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	PrintSlice[string]([]string{"a", "b", "c", "d"}) //型推論によって型を指定しなくてもよい
}
