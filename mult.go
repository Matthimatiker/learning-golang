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
		printNumberCell(i)
		for j := 1; j <= size; j++ {
			printNumberCell(i * j)
		}
		fmt.Println()
	}
}

func printHeadline(size int) {
	printCell(" ")
	for i := 1; i <= size; i++ {
		printNumberCell(i)
	}
	fmt.Println()
}

func printNumberCell(value int) {
	printCell(strconv.Itoa(value))
}

func printCell(value string) {
	fmt.Printf("%4s", value)
	fmt.Print(" ")
}
