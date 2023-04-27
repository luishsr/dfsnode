// Set up a simple HTTP server to handle file uploads and downloads
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Runs the node
func main() {

	// HandleFunc handles upload requests (PUT)
	http.HandleFunc("/upload/", func(w http.ResponseWriter, r *http.Request) {
		fileName := r.URL.Path[len("/upload/"):]
		file, err := os.Create(fileName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		_, err = io.Copy(file, r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	// HandleFunc handles upload requests (GET)
	http.HandleFunc("/download/", func(w http.ResponseWriter, r *http.Request) {
		fileName := r.URL.Path[len("/download/"):]
		file, err := os.Open(fileName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		defer file.Close()

		_, err = io.Copy(w, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Start the HTTP server on port 8000
	port := 8000

	fmt.Printf("Node server starting on port %d...\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		fmt.Println("Failed to start node server:", err)
	}
}
