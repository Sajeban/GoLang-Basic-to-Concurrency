package main

import "fmt"

func main() {
	var name1 string = "Saje"
	var name2 = "Sajeban" //var name string="Sajeban" is also correct but using inverted comma defines it as string
	var age1 int = 21     //In go if you create a variable it must be used or it will throw error
	var age2 = 22
	var isStudent = true
	const pi = 3.14
	name3 := "shorthandmethod" //Shorthand method to define a variable
	fmt.Println(name1, age1, name2, age2, isStudent, name3)
	fmt.Printf("%T\n", pi) //To find the data type of a variable
}
