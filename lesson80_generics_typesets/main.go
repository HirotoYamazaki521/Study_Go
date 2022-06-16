package main

import "fmt"

//generics_typesets
//任意の型の集合を制約として持たせる

type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 //これらいずれかの型を満たしていればNumberのinterfaceを満たしていると判断される

	//~をつけると、素となる方の独自型にも制約を書けることができる
}

func Max[T Number](x, y T) T { //Tの型としてNumberのinterfaceを満たさなければならない
	if x > y {
		return x
	}
	return y
}

type MyInt int

func main() {
	fmt.Println(Max[int](1, 2))
	fmt.Println(Max[float64](1.1, 1.3))
	fmt.Println(Max(1.1, 1.3))

	var x, y MyInt = 1, 2
	fmt.Println(Max[MyInt](x, y))

}
