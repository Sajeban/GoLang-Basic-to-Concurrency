package main

import "fmt"

func main() {
	p := person{name: "Jake", age: 23}
	fmt.Println(p)
	fmt.Println("Age", p.age)

}

//Type is class in Java
type person struct {
	name string
	age  int
}
