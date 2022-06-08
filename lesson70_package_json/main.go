package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

//json
//エンコーダとデコーダの機能を提供

//構造体からJSONテキストへの変換

/*
type A struct{}

type User struct {
	ID int `json:"id"` //何も書かなかったら大文字のIDとなる
	//ID      int      `json:"id,string"` 文字列型に変換
	//ID      int      `json:"id,omitenmpty"` 0や空文字を値を隠すことができる
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       A         `json:"A"`
	//A       *A         `json:"A,omitenmpty"`
}

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "example@example.com"
	u.Created = time.Now()

	//Marshal JSONに変換
	bs, err := json.Marshal(u) //構造体をjsonに変換した値をbyteのスライスで受け取る
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs)) //文字列方に変換して出力

	//----------------------------------------------------
	//Marchalを構造体に戻す

	fmt.Printf("%T\n", bs)

	u2 := new(User)

	//Unmarchal JSONをデータに変換
	if err := json.Unmarshal(bs, &u2); err != nil {
		fmt.Println(err)
	}

	fmt.Println(u2)
}
*/

//Marshalのカスタム
type A struct{}

type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       *A        `json:"A,omitempty"`
}

func (u User) MarshalJSON() ([]byte, error) { //json.Marshal()が呼ばれたときにこのメソッドが呼ばれる
	v, err := json.Marshal(&struct {
		Name string
		//他のフィールドをカスタムしたければここに追加
	}{
		Name: "Mr " + u.Name,
	})
	return v, err
}

func (u *User) UnmarshalJSON(b []byte) error {
	type User2 struct {
		Name string
	}
	var u2 User2
	err := json.Unmarshal(b, &u2)
	if err != nil {
		fmt.Println(err)
	}
	u.Name = u2.Name + "!"
	return err
}

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "example@example.com"
	u.Created = time.Now()

	//Marshal JSONに変換
	bs, err := json.Marshal(u) //構造体をjsonに変換した値をbyteのスライスで受け取る
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs)) //文字列方に変換して出力

	//----------------------------------------------------
	//Marchalを構造体に戻す

	fmt.Printf("%T\n", bs)

	u2 := new(User)

	//Unmarchal JSONをデータに変換
	if err := json.Unmarshal(bs, u2); err != nil {
		fmt.Println(err)
	}

	fmt.Println("u2", u2)
}
