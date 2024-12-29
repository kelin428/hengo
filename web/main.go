package main

import (
	"fmt"
	"hengo"
	"net/http"
)

func main() {
	hen := hengo.New()
	hen.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	hen.Run()
}
