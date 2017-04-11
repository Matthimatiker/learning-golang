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
	printMultiplicationTable(n)
}

func printMultiplicationTable(size int) {
	fmt.Print(" ")
	for i := 1; i <= size; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	for i := 1; i <= size; i++ {
		fmt.Print(i, " ")
		for j := 1; j <= size; j++ {
			fmt.Print(i * j, " ")
		}
		fmt.Println()
	}
}
