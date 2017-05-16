package main

import (
	"net/http"
	"bytes"
	"os"
	"io/ioutil"
)

func main() {
	post("http://tarent.de")
}

func post(url string) {
	res, err := http.Post(url, "text/plain", bytes.NewBufferString("Hello World!"))
	if (err != nil) {
		panic(err)
	}
	res.Header.WriteSubset(os.Stdout, nil)
	body, err := ioutil.ReadAll(res.Body)
	if (err != nil) {
		panic(err)
	}
	println()
	println(string(body))
}
