package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if (err != nil) {
		panic(err)
	}
	mult(n)
}

func mult(n int) {
	fmt.Print(" ")
	for i := 1; i <= n; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	for i := 1; i <= n; i++ {
		fmt.Print(i, " ")
		for j := 1; j <= n; j++ {
			fmt.Print(i * j, " ")
		}
		fmt.Println()
	}
}
