package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{
			Timeout: 5 * time.Second,
		}

		resp2, err := client.Get("http://service2:8081")
		if err != nil {
			log.Printf("Failed to call service2: %v", err)
			http.Error(w, "Failed to call service2", http.StatusInternalServerError)
			return
		}
		defer resp2.Body.Close()
		body2, err := ioutil.ReadAll(resp2.Body)
		if err != nil {
			log.Printf("Failed to read response from service2: %v", err)
			http.Error(w, "Failed to read response from service2", http.StatusInternalServerError)
			return
		}

		resp3, err := client.Get("http://service3:8082")
		if err != nil {
			log.Printf("Failed to call service3: %v", err)
			http.Error(w, "Failed to call service3", http.StatusInternalServerError)
			return
		}
		defer resp3.Body.Close()
		body3, err := ioutil.ReadAll(resp3.Body)
		if err != nil {
			log.Printf("Failed to read response from service3: %v", err)
			http.Error(w, "Failed to read response from service3", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Service1: Calling Service2 and Service3\nService2 Response: %s\nService3 Response: %s", body2, body3)
	})

	log.Println("Service1 is running on port 8080...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
