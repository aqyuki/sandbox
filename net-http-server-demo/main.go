package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /items/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		// logging
		fmt.Printf("Request received for item ID: %s\n", id)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "You specified item ID: %s\n", id)
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Printf("Server is running on %s\n", ":8080")
	server.ListenAndServe()
}
