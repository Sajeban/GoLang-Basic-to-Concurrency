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

}
