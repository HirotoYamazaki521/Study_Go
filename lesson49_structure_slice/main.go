package main

import "fmt"

//構造体　slice

type User struct {
	Name string
	Age  int
	//X, Y int
}

//新しく型を作る
//Users型はUserのポインタ型をスライスとして格納できる
type Users []*User

//下のように定義もできるが上の定義の方法の方が望ましい
/*
type Users struct{
	Users []*Users
}
*/

func main() {
	//複数のUser型をUsersに格納する
	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}
	user4 := User{Name: "user4", Age: 40}

	users := Users{}

	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3, &user4)

	//各データのアドレスが表示
	fmt.Println(users)

	for _, u := range users {
		//uの実体を表示
		fmt.Println(*u)
	}

	//make関数を使って構造体のスライスを生成できる
	users2 := make([]*User, 0)
	users2 = append(users2, &user1, &user2)

	for _, u := range users2 {
		//uの実体を表示
		fmt.Println(*u)
	}

}
