package main

import "testing"

//Debug変数を切り替えて、テストをするかどうかを決める
//false テスト実行
//true テスト実行しない
var Debug bool = false

//goのルールで名前がTestで始まる関数がテスト用
func TestIsOne(t *testing.T) {
	i := 1
	if Debug {
		t.Skip("スキップさせる")
	}
	v := IsOne(i)

	if !v {
		t.Errorf("%v != %v", i, 1)
	}
}

//実行コマンド
/*
go mod init
go mod tidy
go test		test実行
go test -v 詳細表示
*/
