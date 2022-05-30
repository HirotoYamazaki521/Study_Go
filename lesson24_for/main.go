package main

import "fmt"

//for

func main() {
	/*
		i := 0
		for {
			i++
			if i == 3 {
				break
			}
			fmt.Println("Loop")
		}
	*/

	/*
		point := 0
		for point < 10 {
			fmt.Println(point)
			point++
		}
	*/

	/*
		for i := 0; i < 10; i++ {
			if i == 3 {
				continue //for文の中に戻る スキップ
			}
			if i == 6 {
				break
			}
			fmt.Println(i)
		}
	*/

	/*
		arr := [3]int{1, 2, 3}
		for i := 0; i < len(arr); i++ {
			fmt.Println(arr[i])
		}
	*/

	/*
		arr := [3]int{1, 2, 3}

		for i, v := range arr { //1つ目の返り値にインデックス番号、2つ目の返り値に要素の値
			fmt.Println(i, v)
		}
		//キーを使わない場合は、_(アンダーバー)で非表示にできる
	*/

	/*
		sl := []string{"Python", "PHP", "GO"}
		for i, v := range sl {
			fmt.Println(i, v)
		}
	*/

	//map キーバリューの形式でデータを格納
	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}

}
