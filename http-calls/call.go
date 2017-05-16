package main

import (
	"net/http"
	"bytes"
	"os"
	"io/ioutil"
)

func main() {
	res, err := http.Post("http://tarent.de", "text/plain", bytes.NewBufferString("Hello World!"))
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
