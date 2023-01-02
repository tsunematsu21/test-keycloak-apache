package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %v", strings.Join(flag.Args(), ", "))
	})
	log.Fatal(http.ListenAndServe(":18080", nil))
}
