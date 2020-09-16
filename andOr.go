package main

import (
	"fmt"
)

func main() {
	if 6 > 3 && 10 > 5 {
		fmt.Println("True")
	}
	if 5 > 6 || 1 < 5 {
		fmt.Println("True")
	}
}
