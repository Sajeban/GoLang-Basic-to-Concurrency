package main

import "fmt"

func main() {
	var maleactors [2]string   //Declare
	maleactors[0] = "DiCaprio" //Assign
	maleactors[1] = "Brad Pitt"
	fmt.Println(maleactors)

	//Declare and assign
	femaleactors := [2]string{"Margot Robbie", "Sandra Bullock"}
	fmt.Println(femaleactors)

	//Slice
	series := []string{"Big Bang Theory", "Got", "Westworld"}
	fmt.Println(series)
	fmt.Println(len(series))
	fmt.Println(series[1:3])

	//Appending the value at the slice and reassigning to itself
	series = append(series, "BodyGuard")
	fmt.Println(series)

	//Maps  dict-pythons associative arrays-php
	//map(key data type)value datatype
	vertices := make(map[string]int)
	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["rectangle"] = 4
	fmt.Println(vertices)
	//Accessing an element using its key
	fmt.Println(vertices["triangle"])
	//Deleting an element
	delete(vertices, "square")
	fmt.Println(vertices)

}
