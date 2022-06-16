package main

import "fmt"

//generics int

func main() {
	var a any = 1 //go1.18からではinterface型ではなく、any型で指定する方が望ましい
	a = "a"
	a = 1
	a = true
	fmt.Println(a)
}
