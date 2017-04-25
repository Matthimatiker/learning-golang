package main

import (
	"github.com/matthimatiker/learning-golang/util"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "%s", util.Reverse("hello"))
	})
	http.Handle("/", r)
	http.ListenAndServe(":8000", r)
}
