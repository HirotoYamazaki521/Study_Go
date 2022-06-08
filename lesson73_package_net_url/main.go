package main

import (
	"fmt"
	"net/url"
)

//net/url

func main() {
	/*
		//URLを解析
		u, _ := url.Parse("http://example.com/search?a=1&b=2#top")
		fmt.Println(u.Scheme)
		fmt.Println(u.Host)
		fmt.Println(u.Path)
		fmt.Println(u.RawQuery)
		fmt.Println(u.Fragment)

		fmt.Println(u.IsAbs())
		fmt.Println(u.Query())	//mapで取得
	*/

	//URLを生成
	url := &url.URL{} //url型の構造体のポイントを生成
	url.Scheme = "https:"
	url.Host = "google.com"
	q := url.Query()
	q.Set("q", "Go言語")        //"q"がkey、"Golang"がvalue
	url.RawQuery = q.Encode() //クエリーの文字列のエスケープ処理

	fmt.Println(url)
}
