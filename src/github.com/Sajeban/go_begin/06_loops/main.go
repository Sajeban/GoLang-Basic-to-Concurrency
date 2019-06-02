package main

import "fmt"

func main() {
	//Only loop in go is for loop
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println("\nImplementation 2")
	//for loop also can be implemented like while loop
	j := 0
	for j < 5 {
		fmt.Print(j, " ")
		j++
	}
	fmt.Println()
	//Using loops  in maps and arrays
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}

	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"
	for key, value := range m {
		fmt.Println("key", key, "value", value)
	}
}
