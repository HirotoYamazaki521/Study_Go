package main

import "fmt"

//構造体　map

type User struct {
	Name string
	Age  int
	//X, Y int
}

func main() {
	//int型のキー、value User
	m := map[int]User{
		1: {Name: "user1", Age: 10},
		2: {Name: "user2", Age: 20},
	}

	fmt.Println(m)

	//キーとしてUser型を使うこともできる
	m2 := map[User]string{
		{Name: "user1", Age: 10}: "Tokyo",
		{Name: "user2", Age: 20}: "LA",
	}
	fmt.Println(m2)

	//mapもmakeを使って生成できる
	m3 := make(map[int]User)
	fmt.Println(m3)
	m3[1] = User{Name: "user3"}
	fmt.Println(m3)

	for _, v := range m {
		fmt.Println(v)
	}
}
