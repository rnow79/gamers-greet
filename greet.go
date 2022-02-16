package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// HTTP server port.
const port int = 80

// Main function.
func main() {
	// Create router.
	router := mux.NewRouter()
	// Register static directory.
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./front/")))
	// Run web server.
	server := &http.Server{Handler: router, Addr: ":" + strconv.Itoa(port), WriteTimeout: 10 * time.Second, ReadTimeout: 10 * time.Second}
	server.ListenAndServe()
}
