package controllers

import (
	"log"
	"net/http"
	"text/template"
)

//ハンドラー
/*
以下の様な引数をとる関数はハンドラーとして登録できる
*/
func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("app/views/templates/top.html")
	if err != nil {
		log.Fatalln(err)
	}
	/*
		htmlのテンプレートパッケージ
		引数で渡したテンプレートファイルを解析し、
		テンプレートの構造体を生成する
		top.htmlのファイル名をテンプレート名とする新しいテンプレートを生成
	*/
	t.Execute(w, "hello")
	/*
		tのテンプレートファイルを実行
		第1引数はResponseWriter、第2引数はデータを渡す
	*/
}
