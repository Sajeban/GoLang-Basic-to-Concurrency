package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go print("Sajeban", c)
	// for {
	// 	msg, open := <-c
	// 	fmt.Println("main", msg)
	// 	if !open {
	// 		break
	// 	}

	// }
	for msg := range c {
		fmt.Println(msg)
	}
}
func print(x string, c chan int) {
	for i := 0; i <= 5; i++ {
		c <- i
		fmt.Println("routine ", x)
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
