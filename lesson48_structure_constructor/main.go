package main

import "fmt"

//構造体　constructor
//constructor機能はないが関数を使うことはある

type User struct {
	Name string
	Age  int
	//X, Y int
}

//User型を返すコンストラクタ関数を使ってUser型のポインタを生成する
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	user1 := NewUser("user1", 10)

	fmt.Println(user1)

	//user1の実体にアクセス
	fmt.Println(*user1)
}
