package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Service3!")
	})

	log.Println("Service3 is running on port 8082...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8082", nil))
}
