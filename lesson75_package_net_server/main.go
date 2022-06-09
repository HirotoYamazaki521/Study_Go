package main

import (
	"log"
	"net/http"
	"text/template"
)

//net
type MyHandler struct{}

/*
//http.ResponseWriterと*http.Requestのこれらの引数を持つメソッドすべてが
//ハンドラとなる
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")	//ResponseWriterに書き込む
}
*/

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl.html") //テンプレートソースを解析して、テンプレートの構造体を生成
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "Hello World111!") //テンプレートの実行
}

func main() {
	//与えられたURLに応じて各種ハンドラに転送

	//http.ListenAndServe(":8080", &MyHandler{})
	http.HandleFunc("/top", top)      //第1引数にURL、第2引数にハンドラ関数
	http.ListenAndServe(":8080", nil) //ハンドラの引数がnilの場合のマルチプレクサはハンドラを見つけれないため、404
}
