package main

import "fmt"

func main() {
	var saje string
	fmt.Scanln(&saje)

	fmt.Println("Input value :", saje)

	//Getting two integer values

	var a, b int //&a points to the memory address(Pointers)
	fmt.Print("Enter a and b: ")
	fmt.Scanf("%d", &a)

	fmt.Scanf("%d", &b)
	fmt.Println("\nmemory address  of a", &a) //Here we get the memory address
	fmt.Println("Value of a : ", a)           //Here we get the  value
	d := b + a
	fmt.Println("Sum of a & b: ", d)
}
