package main

import (
	"fmt"
	"strings"
)

//strings
//文字列操作の機能

func main() {
	/*
		//文字列を結合する
		s1 := strings.Join([]string{"A", "B", "C"}, ",")
		s2 := strings.Join([]string{"A", "B", "C"}, "")
		fmt.Println(s1, s2)
	*/

	/*
		//文字列に含まれる部分文字列を検索する
		i1 := strings.Index("ABCDE", "A")
		i2 := strings.Index("ABCDE", "D")
		fmt.Println(i1, i2)	//返り値が-1の場合は、部分文字列が含まれていない

		i3 := strings.LastIndex("ABCDABCD", "BC")//最後のインデックス番号を返す
		fmt.Println(i3)

		i4 := strings.IndexAny("ABC", "ABC")	//検索対象の文字列で指定対象の文字のどれかで始まるか
		i5 := strings.LastIndexAny("ABC", "ABC")//最後
		fmt.Println(i4, i5)

		b1 := strings.HasPrefix("ABC", "A")	//検索対象の文字列が指定対象の文字列から始まるか
		b2 := strings.HasSuffix("ABC", "C")	//検索対象の文字列が指定した文字列で終わるか
		fmt.Println(b1, b2)

		b3 := strings.Contains("ABC", "B")	//検索した文字列の中で、指定した文字列が含まれているか
		b4 := strings.ContainsAny("ABCDE", "BD")//検索した文字列の中で、指定した文字列がどちらかが含まれているか
		fmt.Println(b3, b4)

		i6 := strings.Count("ABCABC", "B")
		i7 := strings.Count("ABCABC", "")//文字列の長さ+1
		fmt.Println(i6, i7)
	*/

	/*
		//文字列を繰り返して結合する
		s3 := strings.Repeat("ABC", 4)
		s4 := strings.Repeat("ABC", 0)	//負の値の場合はランタイムエラー
		fmt.Println(s3, s4)
	*/

	/*
		//文字列の置換
		s5 := strings.Replace("AAAAAA", "A", "B", 1) //1は置換する回数の最大数
		s6 := strings.Replace("AAAAAA", "A", "B", -1)//-1は該当するすべての箇所
		fmt.Println(s5, s6)
	*/

	/*
		//文字列を分割する
		s7 := strings.Split("A,B,C,D,E", ",")
		fmt.Println(s7)
		s8 := strings.SplitAfter("A,B,C,D,E", ",")
		fmt.Println(s8)
		s9 := strings.SplitN("A,B,C,D,E", ",", 2)
		fmt.Println(s9)
		s10 := strings.SplitAfterN("A,B,C,D,E", ",", 4)
		fmt.Println(s10)
	*/

	/*
		//大文字小文字の返還
		s11 := strings.ToLower("ABC")
		s12 := strings.ToLower("E")

		s13 := strings.ToUpper("abc")
		s14 := strings.ToUpper("e")
		fmt.Println(s11, s12, s13, s14)
	*/

	/*
		//文字列からスペースで区切られたフィールドを取り出す
		s15 := strings.TrimSpace("    -    ABC    -    ")
		//全角
		s16 := strings.TrimSpace("　　　　ABC　　　　")
		fmt.Println(s15, s16)
	*/

	//文字列からスペースで区切られたフィールドを取り出す
	s18 := strings.Fields("a b c") //返り値は文字列型のスライス
	fmt.Println(s18)
	fmt.Println(s18[1])

}
