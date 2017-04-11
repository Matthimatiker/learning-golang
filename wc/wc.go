package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	words := 0
	for scanner.Scan() {
		words++
	}
	fmt.Println(words)
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
