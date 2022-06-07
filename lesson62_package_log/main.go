package main

import (
	"log"
	"os"
)

//log

func main() {
	//ログの出力先を変更
	//log.SetOutput(os.Stdout)

	/*
		log.Print("Log\n") //デフォルトで日付と時刻
		log.Println("Log2")
		log.Printf("Log%d\n", 3)
	*/

	/*
		//opのexitを伴うため、log.Fatalのみ出力
		log.Fatal("Log\n")
		log.Fatalln("Log2")
		log.Fatalf("Log%d\n", 3)
	*/

	/*
		//panicは文字列を出力してpanic状態にし、プログラムを強制終了
		log.Panic("Log\n")
		log.Panicln("Log2")
		log.Panicf("Log%d\n", 3)
	*/

	/*
		//任意のファイルを作成し、出力先に指定
		//os.Create ファイルの作成
		f, err := os.Create("test.log")
		if err != nil {
			return
		}
		//作成したio.Writer型のファイルを出力先に設定
		log.SetOutput(f)
		log.Println("ファイルに書き込む")
	*/

	//ログのフォーマットを指定する
	/*
		//標準のログフォーマット
		log.SetFlags(log.LstdFlags)
		log.Println("A")

		//マイクロ秒を追加
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.Println("B")

		//時刻とファイルの行番号(短縮系)
		log.SetFlags(log.Ltime | log.Lshortfile)
		log.Println("C")

		log.SetFlags(log.LstdFlags)

		//ログのプリフィックスを設定
		log.SetPrefix("[LOG]")
		log.Println("E")
	*/

	//ロガーの生成
	//ログパッケージの関数はデフォルトで設定されている１つのロガーを全体に適用しているので、
	//一部のログの出力を変えたい場合など、小回りが利かない
	//ロガーを新規で生成する
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("message")
	log.Println("message")

	_, err := os.Open("fdafdsafa")
	//ログはエラーハンドリングで使われることが多い
	if err != nil {
		//ログ出力
		//log.Fatalln("Exit", err)
		logger.Fatalln("Exit", err)
	}

}
