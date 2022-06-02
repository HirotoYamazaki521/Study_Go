package main

import "fmt"

//structure メソッド

type User struct {
	Name string
	Age  int
	//X, Y int
}

//メソッドの定義
//レシーバーの宣言
func (u User) SayName() {
	fmt.Println(u.Name)
}

//ポインタ型でレシーバーを参照渡ししたいとき
func (u User) SetName(name string) {
	u.Name = name
}

//基本的にメソッドのレシーバはポインタ型が望ましい
func (u *User) SetName2(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "user1"}
	user1.SayName()

	user1.SetName("A")
	user1.SayName()

	user1.SetName2("A")
	//呼び出し側はポインタ型にする必要がない
	user1.SayName()

	user2 := &User{Name: "user2"}
	user2.SetName2("B")
	user2.SayName()

}
