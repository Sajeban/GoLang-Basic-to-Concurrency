package main

import "fmt"

func main() {
	i := 7
	inc(i) //We are passing a value so no change will be done to our original variable
	fmt.Println(i)
	//Through Pointers we can access the original variable through its memory address

	increase(&i)
	fmt.Println("After accessing through memory")
	fmt.Println(i)

}
func inc(x int) {
	x++
}
func increase(x *int) {
	*x++
	//x++ will increase the memory address put an astrick before the variable
}
