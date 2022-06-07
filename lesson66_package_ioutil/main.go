package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//ioutil
//簡易的なファイルの読み書き

func main() {
	//入力全体を読み込む
	f, _ := os.Open("foo.txt")
	bs, _ := ioutil.ReadAll(f) //byteのスライスに変換
	fmt.Println(string(bs))

	//ファイルに書き込み
	if err := ioutil.WriteFile("bar.txt", bs, 0666); err != nil {
		log.Fatalln(err)
	}
}
