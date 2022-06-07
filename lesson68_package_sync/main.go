package main

import (
	"fmt"
	"sync"
)

//sync

//ミューテックスによる同期処理
//レースコンディションを防ぐ
//レースコンディションとは複数の処理が同じデータに同時にアクセスした場合に、機能停止など予期しない処理結果が生じてしまうこと

//同期処理の失敗例
/*
var st struct{ A, B, C int }

func UpdateAndPrint(n int) {
	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st)
}

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		}()
	}
	for {

	}
}
*/

//ミューテックスによる同期処理
/*
var st struct{ A, B, C int }

//ミューテックスを保持するパッケージ変数
var mutex *sync.Mutex

func UpdateAndPrint(n int) {
	//ロック
	//1つのゴルーチンのみがロックを取得できる性質があり、
	//ロックされている間は、1つのゴルーチンしか処理できない
	mutex.Lock()

	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st)

	//アンロック
	mutex.Unlock()
}

func main() {
	mutex = new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		}()
	}
	for {
	}
}
*/

//ゴルーチンの処理を待ち受ける

func main() {
	/*
		go func() {
			for i := 0; i < 100; i++ {
				fmt.Println("1st Goroutine")
			}
		}()
		go func() {
			for i := 0; i < 100; i++ {
				fmt.Println("2nd Goroutine")
			}
		}()
		go func() {
			for i := 0; i < 100; i++ {
				fmt.Println("3rd Goroutine")
			}
		}()

		//無限ループがないと、非同期処理が実行されることなく終了する
		//for {}
	*/
	//sunc.WaitGroupを生成
	wg := new(sync.WaitGroup)
	// 待ち受けするゴルーチンの数は3
	wg.Add(3)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
		wg.Done() //完了
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Goroutine")
		}
		wg.Done() //完了
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Goroutine")
		}
		wg.Done() //完了
	}()

	//ゴルーチンの完了を待ち受ける
	//Doneが3つ完了するまで待つ

	wg.Wait()

}
