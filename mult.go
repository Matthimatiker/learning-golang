package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Exercise from https://github.com/smancke/talks/blob/gh-pages/golang_schulung/01_basics.md#%C3%9Cbung-2-multiplikationstabelle
 */
func main() {
	n, err := strconv.Atoi(os.Args[1])
	if (err != nil) {
		panic(err)
	}
	printMultiplicationTable(n)
}

func printMultiplicationTable(size int) {
	printHeadline(size)
	for i := 1; i <= size; i++ {
		fmt.Print(i, " ")
		for j := 1; j <= size; j++ {
			fmt.Print(i * j, " ")
		}
		fmt.Println()
	}
}

func printHeadline(size int) {
	fmt.Print(" ")
	for i := 1; i <= size; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
