package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/matthimatiker/learning-golang/util"
)

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := ReverseReader.NewReverseReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(reverse(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func reverse(bytes string) string {
	strLength := len(bytes)
	reversed := make([]byte, strLength, strLength)
	for i := 0; i < strLength; i++ {
		reversed[i] = bytes[strLength-i-1]
	}
	return string(reversed)
}
