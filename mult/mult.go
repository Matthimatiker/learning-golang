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
	padding := calcPadding(size)
	printHeadline(size, padding)
	for i := 1; i <= size; i++ {
		printNumberCell(i, padding)
		for j := 1; j <= size; j++ {
			printNumberCell(i * j, padding)
		}
		fmt.Println()
	}
}

func printHeadline(size int, padding int) {
	printCell(" ", padding)
	for i := 1; i <= size; i++ {
		printNumberCell(i, padding)
	}
	fmt.Println()
}

func printNumberCell(value int, padding int) {
	printCell(strconv.Itoa(value), padding)
}

func printCell(value string, padding int) {
	fmt.Printf("%" + strconv.Itoa(padding) + "s", value)
	fmt.Print(" ")
}

func calcPadding(size int) (int) {
	return len(strconv.Itoa(size * size))
}
