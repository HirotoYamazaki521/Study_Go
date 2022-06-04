package main

import (
	"fmt"
	"log"
	"os"
)

//os

func main() {
	//Exit
	/*
		os.Exit(1)
		fmt.Println("Start")

		defer func() {
			fmt.Println("defer")
		}()
		os.Exit(0)
	*/

	//log.Fatal
	/*
		_, err := os.Open("A.txt")
		if err != nil {
			log.Fatalln(err)
		}
	*/

	//Args
	/*
		//String型のスライスで任意のコマンドライン引数が格納されている
		fmt.Println(os.Args[0])
		fmt.Println(os.Args[1])
		fmt.Println(os.Args[2])
		fmt.Println(os.Args[3])

		//os.Argsの要素数を表示
		fmt.Println("length=%d\n", len(os.Args))
		//os.Argsの中身をすべて表示
		for i, v := range os.Args {
			fmt.Println(i, v)
		}

		//実行ファイルの作成
		//go build -o getargs(実行ファイル名) main.go

		//実行ファイルの実行
		//./getargs 123 456 789

	*/

	//ファイル操作
	//Open
	/*
		f, err := os.Open("test.txt")
		if err != nil {
			log.Fatalln(err)
		}

		defer f.Close()
	*/

	//ファイル操作
	//Create
	/*
		//指定したファイルが存在していた場合は、削除されて新規で作成される
		f, _ := os.Create("foo.txt")

		f.Write([]byte("Hello\n"))

		f.WriteAt([]byte("Golang"), 6)	//オフセット1を6文字目にセット

		f.Seek(0, os.SEEK_END)	//ファイルの末尾にオフセットを移動

		f.WriteString("Yaah")
	*/

	//ファイル操作
	//Read
	/*
		f, err := os.Open("foo.txt")
		if err != nil {
			log.Fatalln(err)
		}

		defer f.Close()

		bs := make([]byte, 128)

		n, err := f.Read(bs)	//byteのスライスに読み込んだ内容を書き込み
		fmt.Println(n)
		fmt.Println(string(bs))	//byteのスライスを文字列に変換して表示

		bs2 := make([]byte, 128)

		nn, err := f.ReadAt(bs2, 10)//第1引数にスライス、第2引数にオフセット
		fmt.Println(nn)
		fmt.Println(string(bs2))
	*/

	//OpenFile
	//O_RDONLY 読み込み専用
	//O_WRONLY 書き込み専用
	//O_RDWR 読み書き専用
	//O_APPEND ファイル末尾に追記
	//O_CREATE ファイルがなければ作成
	//O_TRUNC 可能であればファイルの内容をオープン時に空にする

	f, err := os.OpenFile("foo.txt", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))

}
