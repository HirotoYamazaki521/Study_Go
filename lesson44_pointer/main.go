package main

import "fmt"

//ポインタ
//値型に分類されるデータ構造

func Double(i int) {
	i = i * 2
}

func Doublev2(i *int) {
	*i = *i * 2
}

func Doublev3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	var n int = 100
	fmt.Println(n)

	//メモリのアドレスを表示
	fmt.Println(&n)

	Double(n)
	//nからコピーされた値がDouble関数で2倍されてコピーされたiが2倍されているので
	//もとのnの値は変わることなく100のまま

	fmt.Println(n)

	var p *int = &n //*データ型でポインタ型指定できる　&はアドレス演算子
	fmt.Println(p)

	//デリファレンス
	//ポインタ型が保持するメモリ上のアドレスを経由してデータ本体を参照する仕組み
	fmt.Println(*p)

	/*
		*p = 300
		fmt.Println(n)
		n = 200
		fmt.Println(*p)
	*/

	Doublev2(&n) //参照渡し
	fmt.Println(n)

	Doublev2(p)
	fmt.Println(*p)

	//参照型はもともと参照渡し機能を持っているので
	//ポインタ型を使わなくても値を書き換えることができる
	var sl []int = []int{1, 2, 3}

	Doublev3(sl)
	fmt.Println(sl)
}
