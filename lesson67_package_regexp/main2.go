//package main

import (
	"fmt"
	"regexp"
)

//regexp

func main() {

	/*
		re := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)

		//正規表現にマッチした文字列の取得
		//FindString
		fs1 := re.FindString("AAAAABCXYZZZZ")	//初めにマッチした部分をstringで返す
		fmt.Println(fs1)
		fs2 := re.FindAllString("ABCXYZABCXYZ", -1)//-1 マッチしたすべての文字列をスライス型で返す
		fmt.Println(fs2)
	*/

	//正規表現のグループによるサブマッチ
	//FindAllStinrgSubmatch
	re1 := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`) //\d*は1回以上繰り返す整数値
	s := `
	0123-456-7889
	111-222-333
	556-787-899
	`

	ms := re1.FindAllStringSubmatch(s, -1) //マッチしたすべての要素を取得
	fmt.Println(ms)                        //スライスで[マッチした要素 サブマッチした要素]と格納される

	for _, v := range ms {
		fmt.Println(v)
	}

	fmt.Println(ms[0][0])
	fmt.Println(ms[0][1])
	fmt.Println(ms[0][2])

}
