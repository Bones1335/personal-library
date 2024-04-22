package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>This is the %q page</h1>\n", r.URL.Path)
}
