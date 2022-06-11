package controllers

import (
	"lesson77_Todo_app/config"
	"net/http"
)

func StartMainServer() error {
	//css、jsファイルを読み込む
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	/*
		staticの改装の下にcss、jsファイルが必要になってくるがstaticフォルダがないので、
		http.StripPrefix("/static/", files)でstaticを取り除いてfilesを指定

	*/

	//URLを登録してhtmlファイルを返して表示する
	http.HandleFunc("/", top) // 第2引数にハンドラー
	/*
		URLを登録しており、

	*/
	return http.ListenAndServe(":"+config.Config.Port, nil)
	/*
		config.iniで設定したポート番号の値を読み込む
		第2引数は通常、デフォルトのマルチプレクサを使う
		デフォルトのマルチプレクサでは登録されていないURLにアクセスしたら
		404 page not found を返す
	*/

}
