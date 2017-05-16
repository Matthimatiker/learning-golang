package main

import (
	"net/http"
	"bytes"
	"io/ioutil"
)

func main() {
	println(post("http://tarent.de"))
}

func post(url string) string {
	res, err := http.Post(url, "text/plain", bytes.NewBufferString("Hello World!"))
	if (err != nil) {
		panic(err)
	}
	result := bytes.NewBufferString("")
	res.Header.WriteSubset(result, nil)
	body, err := ioutil.ReadAll(res.Body)
	if (err != nil) {
		panic(err)
	}
	result.WriteString("\n")
	result.WriteString(string(body))
	return result.String()
}
