package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "dev"
	}
	fmt.Fprintf(w, "Hello from Argo Rollouts Canary Demo! Version: %s\n", version)
}

func main() {
	http.HandleFunc("/", handler)
	port := "8080"
	fmt.Println("Server running on port", port)
	http.ListenAndServe(":"+port, nil)
}
