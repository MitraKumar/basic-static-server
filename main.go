package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/", fs)

	fmt.Printf("Listening on PORT - :8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Something Went wrong...")
	}
}
