package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s request received: %s:%s:%s\n", r.Method, r.UserAgent(), r.URL, r.RemoteAddr)
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Error getting server details", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "This request is being served by sever %s\n", hostname)
}

func main() {

	port := 3000

	if sp := os.Getenv("SERVER_PORT"); sp != "" {
		p, err := strconv.Atoi(sp)
		if err != nil {
			log.Fatal("Invalid port provided by env var SERVER_PORT: ", sp)
		}
		port = p
	}

	http.HandleFunc("/", rootHandler)

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting server on address: %s\n\n", addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}
