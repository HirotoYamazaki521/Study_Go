package main

import "fmt"

//構造体
//複数の任意の型の値を1つにまとめたもの

//構造体の定義
type User struct {
	Name string //Nameフィールドstrign型
	Age  int    //Ageフィールド int型
	//X,Y int	まとめて定義も可能
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	//明示的にUser型を定義
	var user1 User
	fmt.Println(user1) //{ 0} 1つ目のデータはstrign型の初期値の空文字、2つ目はint型の初期値の0

	//フィールドを更新
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	//暗黙的にUser型を定義
	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	fmt.Println(user2)

	//変数宣言時にフィールド値を指定
	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	user4 := User{"user4", 40} //フィールドを指定しないで、値を指定する場合はフィールドの順番にする必要あり
	fmt.Println(user4)

	//user5 := User{30, "user5"}
	//fmt.Println(user5)

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	user7 := new(User) //newで定義したUserは構造体のポインタ型を返す
	fmt.Println(user7)

	//ポインタ型で指定
	//使う場面としては関数の引数として構造体の変数を渡す場合にポインタ型の時
	user8 := &User{}
	fmt.Println(user8)

	UpdateUser(user1)
	UpdateUser2(user8)

	fmt.Println(user1) //更新されていない
	fmt.Println(user8) //参照渡しできている

}
