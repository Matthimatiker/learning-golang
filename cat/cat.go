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

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(util.Identity(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
