package foo

//メインパッケージとは分割されたパッケージ

const (
	Max = 100
	min = 1 //識別子の一文字目が小文字なのでprivate、外部パッケージから参照不可
)

func ReturnMin() int { //publicの関数
	return min
}
