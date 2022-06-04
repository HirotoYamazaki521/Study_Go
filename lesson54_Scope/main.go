package main

import (
	"Study_Go/lesson54_Scope/foo"
	"fmt"
	. "fmt" //あまり推奨されていない
	f "fmt" //任意のパッケージ名を付与
	//アルファベット順に並び替えると可読性が上がる
	//foo.packageで"time"packageをいれても、main.goで使うことはできないので、
	//パッケージ事にインポートしなければならない
)

//スコープ

//プログラムコードの任意の場所でどのような変数、定数、
//関数が参照可能かはすべてスコープによって決定される

//関数のスコープ
func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	//返り値と変数が同名なので、変数の再定義となりエラーが発生
	//var b string = s
	b = s
	//関数内で重複して識別子を使いたい場合
	{ //ブロック
		b := "BBBB"
		fmt.Println(b)
	}
	fmt.Println(b)

	return b
}

func main() {
	fmt.Println(foo.Max)
	//fmt.Println(foo.min)

	f.Println(foo.ReturnMin()) //関数の頭文字が大文字だから呼び出せる
	Println(foo.Max)

	fmt.Println(appName())

	//関数の中で定義された変数は関数の中でのみ参照可能
	//fmt.Println(AppName,Version)

	fmt.Println(Do("AAA"))
}
