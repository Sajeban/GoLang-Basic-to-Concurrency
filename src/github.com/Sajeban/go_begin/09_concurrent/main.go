package main

import (
	"fmt"
	"time"
)

func main() {
	go print("Sajeban")
	 print("Vaksh")
}
func print(x string) {
	for i := 0; true; i++ {
		fmt.Println(x)
		time.Sleep(time.Millisecond * 500)
	}
}