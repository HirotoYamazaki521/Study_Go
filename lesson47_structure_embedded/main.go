package main

import "fmt"

//構造体 埋め込み

//User型をフィールドとしてTの構造体を埋め込む
type T struct {
	User User

	//型名を省略可能
	//User
}

type User struct {
	Name string
	Age  int
	//X, Y int
}

func (u *User) SetName() {
	u.Name = "A"
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}
	fmt.Println(t)

	//内側のUserフィールドにアクセス
	fmt.Println(t.User)

	fmt.Println(t.User.Name)

	//型名を省略した場合、User型のフィールドに直接アクセス可能
	//fmt.Println(t.Name)

	t.User.SetName()
	fmt.Println(t)
}
