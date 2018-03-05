package main

import "fmt"

func main() {
	var a int
	a = MyFunc()
	fmt.Println("a =", a)

	b := MyFunc()
	fmt.Println("b =", b)
}

//MyFunc MyFunc
func MyFunc() int {
	return 777
}
