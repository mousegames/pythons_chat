package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

var content string = "you are the first one here"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			text, _ := ioutil.ReadAll(r.Body)
			content = string(text)
		} else {
			fmt.Fprintf(w, "%s", content)
		}
	})
	http.ListenAndServe(":8000", nil)
}