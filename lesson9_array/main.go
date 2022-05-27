package main

import "fmt"

//配列型
//他の言語と違い、後から要素数を変更することができない
//スライス型は要素数を変更可能

func main() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1) //"[3]int"ここまでが型となるため、サイズの変更ができない

	var arr2 [3]string = [3]string{"A", "B"} //3つ目に文字列型の空文字が入る
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	//要素数の省略 要素数を自動でカウント
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	//値の取り出し
	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])
	fmt.Println(arr1[2-1])

	//値の更新
	arr2[2] = "C"
	fmt.Println(arr2)

	/*
		var arr5 [4]int
		arr5 = arr1
		fmt.Println(arr5)
		中の要素が同じでも要素数が違えば、違う型として扱われる
	*/

	//len関数で中の要素数を調べることができる
	fmt.Println(len(arr1))
}
