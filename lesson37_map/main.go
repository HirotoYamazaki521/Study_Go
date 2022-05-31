package main

import "fmt"

//マップ
//Pythonでいう辞書型というデータ構造
//キーバリューの形式で配列を作ることができる

func main() {
	//明示的宣言
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	//暗黙的宣言
	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	//make関数でmapを宣言
	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

	//値の取り出し
	fmt.Println(m["A"])
	fmt.Println(m4[2])
	fmt.Println(m4[3]) //登録されていない値を取得すると初期値(空文字)

	//上記のように空文字を取得しても処理が続行されてしまう
	//そのため、mapにはエラーハンドリング機能もついている
	s, ok := m4[3]
	if !ok { //okがfalseだったらエラーを表示
		fmt.Println("error")
	}
	fmt.Println(s, ok) //2つ目の返り値には取り出しに成功したかの判定が帰ってくる

	//値の更新
	m4[2] = "US"
	fmt.Println(m4)

	m4[3] = "CHINA"
	fmt.Println(m4)

	//要素の削除
	delete(m4, 3) //第1引数に削除したいmap,第2引数に削除したいキー
	fmt.Println(m4)

	//mapのlen関数
	fmt.Println(len(m4))

}
