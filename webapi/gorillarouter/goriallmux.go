package gorillamux

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	// mux returns all path paramters as a map
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Genre is: %v\n", vars["genre"])
	fmt.Fprintf(w, "Rating is: %v\n", vars["rated"])
}

func DirectorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is Directors path: %v\n", mux.Vars(r)["id"])
}

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Details path: %v\n", mux.Vars(r)["id"])
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Query Parameter: %s", queryParams["title"])
	fmt.Fprintf(w, "Query Parameter: %s", queryParams["genre"])
}

func RunGorillaMux() {
	r := mux.NewRouter()

	// // Creating SubRouter
	s := r.PathPrefix("/movies").Subrouter()
	s.HandleFunc("{id}/directors", DirectorHandler)
	s.HandleFunc("{id}/details", DetailsHandler)

	r.HandleFunc("/movies", QueryHandler)
	r.Queries("title","genre")

	// Attach a path with Handler
	r.HandleFunc("/movies/{genre}/{rated}", ArticleHandler)
	srv := &http.Server{
		Handler: r,
		Addr:    ":9000",
	}

	log.Fatal(srv.ListenAndServe())
}
