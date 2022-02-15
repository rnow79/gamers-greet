package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// HTTP server port.
const port int = 80

// Redirect function.
func greetRedirection(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/front/index.htm", http.StatusTemporaryRedirect)
}

// Main function.
func main() {
	// Create router.
	router := mux.NewRouter()
	// Register static directory.
	router.PathPrefix("/front").Handler(http.StripPrefix("/front/", http.FileServer(http.Dir("./front/"))))
	// Register redirection.
	router.HandleFunc("/", greetRedirection).Methods(http.MethodGet)
	// Run web server.
	server := &http.Server{Handler: router, Addr: ":" + strconv.Itoa(port), WriteTimeout: 10 * time.Second, ReadTimeout: 10 * time.Second}
	server.ListenAndServe()
}
