package main

import (
	"Study_Go/lesson55_test/alib"
	"fmt"
)

//テスト

//goの標準パッケージにtestパッケージがある
//テストしたい階層と同じ改装ファイルを置く

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsOne(1))
	fmt.Println(IsOne(0))

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(alib.Average(s))
}

//実行コマンド
/*
go mod init
go mod tidy
go test -v ./... ディレクトリ内のすべてのパッケージをテスト
go test -cover ./... パッケージの変数のテストのカバー率が表示 すべての関数をテストする必要あり
*/
