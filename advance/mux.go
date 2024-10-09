package advance

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HomeHandler handles the "/" route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

// ArticlesHandler handles the "/articles" route
func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is the Articles page.")
}

// ArticlesCategoryHandler handles articles with category, demonstrating how to use URL variables
func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

// Mux sets up the router and starts the HTTP server
func Mux() {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/articles", ArticlesHandler)
	// Example of a route with a URL variable
	r.HandleFunc("/articles/{category}", ArticlesCategoryHandler)

	// Start the server
	http.Handle("/", r)
	fmt.Println("Starting server on :3000")

	// // this is also good but we will use log package
	// err := http.ListenAndServe(":3000", nil)
	// if err != nil {
	// 	fmt.Println("Error starting server:", err)
	// }

	log.Fatal(http.ListenAndServe(":3000", nil))

}
