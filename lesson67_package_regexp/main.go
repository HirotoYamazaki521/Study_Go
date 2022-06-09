package main_

import (
	"fmt"
	"regexp"
)

//regexp
//正規表現による文字列処理のためのパッケージ

func main() {

	//Goの正規表現の基本 regexp.MatchString
	match, _ := regexp.MatchString("A", "ABC")
	fmt.Println(match)

	/*
		//Compile
		re1, _ := regexp.Compile("A") //第1引数に文字、第2引数にエラー
		match = re1.MatchString("ABC")
		fmt.Println(match)

		//MustCompile　ランタイムエラーが発生
		re2 := regexp.MustCompile("A")
		match = re2.MatchString("ABC")
		fmt.Println(match)

		//regexp.MustCompile("\\d")
		//regexp.MustCompile(`\d`)
	*/

	//正規表現のフラグ
	/*
		フラグ一覧

		i 大文字小文字を区別しない
		m マルチラインモード(^と&が文頭,文末に加えて行頭、行末にマッチ)
		s .が\nにマッチ
		u 最小マッチへの変換(x*はx*?へ、x+はx+?へ)
	*/
	re3 := regexp.MustCompile(`(?i)abc`) //(?i-ms)iを有効化して、msを無効化する
	match = re3.MatchString("ABC")
	fmt.Println(match)

	//幅をもたない正規表現のパターン
	/*
		パターン一覧

		^ 文頭(mフラグが有効な場合は行頭にも)
		$ 文末(mフラグが有効な場合は行末にも)
		\A 文頭
		\z 文末
		\b ASCIIによるワード協会
		\B 非ASCIIによるワード協会
	*/

	re4 := regexp.MustCompile(`^ABC$`)
	match = re4.MatchString("ABC")

	//繰り返しを表す正規表現
	/*
		繰り返しのパターン

		x* 0回以上繰り返すx(最大マッチ)
		x+ 1回以上繰り返すx(最大マッチ)
		x? 0回以上1回以下繰り返すx
		x{n,m} n回以上m回以下繰り返すx(最大マッチ)
		x{n, } n回以上繰り返すx(最大マッチ)
		x{n} n回繰り返すx(最大マッチ)
		x*? 0回以上繰り返すx(最小マッチ)
		x+? 1回以上繰り返すx(最小マッチ)
		x?? 0回以上1回以下繰り返すx(0回優先)
		x{n,m}? n回以上m回以下繰り返すx(最小マッチ)
		x{n, }? n回以上繰り返すx(最小マッチ)
		x{n}? n回繰り返すx(最小マッチ)
	*/
	re6 := regexp.MustCompile("a+b*")
	fmt.Println(re6.MatchString("ab"))
	fmt.Println(re6.MatchString("a"))
	fmt.Println(re6.MatchString("aaaabbbbbbbb"))
	fmt.Println(re6.MatchString("b"))

	//正規表現の文字クラス
	re8 := regexp.MustCompile(`[XYZ]`) //XYZのどれかにマッチする
	fmt.Println(re8.MatchString("Y"))

	re9 := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`) //3文字
	fmt.Println(re9.MatchString("ABC"))
	fmt.Println(re9.MatchString("abcdefg"))

	re10 := regexp.MustCompile(`[^0-9A-Za-z_]`) //以外
	fmt.Println(re10.MatchString("ABC"))
	fmt.Println(re10.MatchString("あ"))

	//正規表現のグループ
	/*
		(正規表現) グループ (順序によるキャプチャ)
		(?:正規表現) グループ (キャプチャされない)
		(?:P<name> 正規表現) 名前付きグループ
	*/

	re11 := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
	fmt.Println(re11.MatchString("abcxyz"))
	fmt.Println(re11.MatchString("ABCXYZ"))
	fmt.Println(re11.MatchString("ABCxyz"))
	fmt.Println(re11.MatchString("ABCabc"))
}
