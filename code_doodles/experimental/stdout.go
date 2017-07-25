package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(time.Second)
	fmt.Println("\nBlow me")
	for i := 1; i <= 10; i++ {
		<-ticker
		// x0c
		if i%2 == 0 {
			fmt.Println("\r akio")
		}
		fmt.Printf("\rOn %d/10", i)
	}
	fmt.Println("\nAll is said and done.")
}
