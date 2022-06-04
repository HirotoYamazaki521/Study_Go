package main

import (
	"fmt"
	"time"
)

//time

func main() {
	//時間の生成
	/*
		//今の時間
		t := time.Now()
		fmt.Println(t)

		//指定した時間を生成
		t2 := time.Date(2020, 6, 10, 10, 10, 10, 0, time.Local)
		fmt.Println(t2)
		fmt.Println(t2.Year())
		fmt.Println(t2.YearDay())
		fmt.Println(t2.Month())
		fmt.Println(t2.Weekday())
		fmt.Println(t2.Day())
		fmt.Println(t2.Hour())
		fmt.Println(t2.Minute())
		fmt.Println(t2.Second())
		fmt.Println(t2.Nanosecond())
		fmt.Println(t2.Zone())
	*/

	//時刻の感覚を表現する
	/*
		//time.Duration型を返す
		fmt.Println(time.Hour)
		fmt.Printf("%T\n", time.Hour)
		fmt.Println(time.Minute)
		fmt.Println(time.Second)
		fmt.Println(time.Millisecond)
		fmt.Println(time.Microsecond)
		fmt.Println(time.Nanosecond)

		//time.Duration型を文字列から生成する
		d, _ := time.ParseDuration("2h30m")	//返り値に値とエラーを受け取る
		fmt.Println(d)

		t3 := time.Now()
		t3 = t3.Add(2*time.Minute + 15*time.Second)	//2分と15秒を追加
		fmt.Println(t3)
		//現在時刻の2分15秒後を表すtime.Time型の取得
	*/

	//時刻の差分を取得
	/*
		t5 := time.Date(2020, 7, 24, 0, 0, 0, 0, time.Local)
		t6 := time.Now()

		//t5-t6
		d2 := t5.Sub(t6) //メソッドSubで差分を取得
		fmt.Println(d2)

		//時刻を比較する
		fmt.Println(t6.Before(t5))	//t6がt5より前か
		fmt.Println(t6.After(t5))
		fmt.Println(t5.Before(t6))
		fmt.Println(t5.After(t6))
	*/

	//指定時間のスリープ
	//5秒間停止
	time.Sleep(5 * time.Second)
	fmt.Println("5秒間停止後表示")
}
