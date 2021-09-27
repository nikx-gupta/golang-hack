package middleware

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func logMainHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing Main Handler")
	w.Write([]byte("OK"))
	log.Println("Request Completed")
}

func RunWithLoggingMiddleware() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler)
	logMiddleware := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":9000", logMiddleware)
}
