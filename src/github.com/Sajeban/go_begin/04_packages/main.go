package main

import "fmt"

func main() {
	fmt.Println(sayHello("Saje"))
	fmt.Println(add(4, 3))
	fmt.Println(subtract(4, 3))
}

func sayHello(name string) string {
	return "Hello" + name
}
func add(num1 int, num2 int) int {
	return num1 + num2
}
func subtract(num1, num2 int) int { //if both parameters are same we can use this way too
	return num1 - num2
}
