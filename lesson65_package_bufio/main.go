package main

import (
	"bufio"
	"fmt"
	"os"
)

//bufio
//goの基本的な入出力処理にバッファ処理の機能をまとめたもの

func main() {
	//標準入力を行単位で読み込む

	//標準入力をソースにしたスキャナの生成
	//ioReader型の引数を入力もとにしたbufioスキャナー型を生成する
	//改行を区切りとしたスキャン処理を行う
	scanner := bufio.NewScanner(os.Stdin)

	//入力のスキャンが成功する限り繰り返すループ
	for scanner.Scan() { //Scan処理が成功する限りtrueを返し続ける
		//スキャン内容を文字列で出力
		fmt.Println(scanner.Text())
	}

	//スキャンにエラーが発生した場合の処理
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー:", err)
	}
}
