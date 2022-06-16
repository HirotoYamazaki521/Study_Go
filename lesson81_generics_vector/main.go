package main

import "fmt"

//generics_vector
//型定義

type Vector[T any] []T

type IntVector = Vector[int] //インスタンス化

func main() {
	var v Vector[int] = Vector[int]{10, 20, 30}
	fmt.Println(v)

	var v2 Vector[float64] = Vector[float64]{1.3, 3.4, 5.6}
	fmt.Println(v2)

	v3 := IntVector{1, 2, 3}
	fmt.Println(v3)

}
