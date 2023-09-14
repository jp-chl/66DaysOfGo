// server.go

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/get_something", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		response := processingRequest(name)
		fmt.Fprint(w, response)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func processingRequest(name string) string {
	fmt.Println("[DEBUG] processing request..")
	return "Hi there! You requested " + name
}
