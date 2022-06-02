package main

import "fmt"

//interface
//カスタムエラー

//interfaceで独自のカスタムエラーを定義し作ることができる
/*
type error interface{
	Error() string
}
*/

type MyError struct {
	Message string
	ErrCode int
}

//goのソースコードのErrorというinterfaceの型と共通の性質持つ型として
//MyErrorが認識される
func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 1234}
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error())

	//変数のerrはあくまでプログラム上の変数なのでMyErrorに定義されているフィールドにアクセスすることができない
	//fmt.Println(err.Message)

	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}
}
