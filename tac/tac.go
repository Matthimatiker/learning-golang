package main

import (
	//"bufio"
	"fmt"
	"os"
	"bufio"
)

func main() {
	file := "tac/tac.go"
	fmt.Println(file)
	filePointer, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(filePointer)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
