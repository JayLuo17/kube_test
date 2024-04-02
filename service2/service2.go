package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Service2!")
	})

	log.Println("Service2 is running on port 8081...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil))
}
