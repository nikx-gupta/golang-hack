package middleware

import (
	"fmt"
	"net/http"
)

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Executing middleware before request phase")
		handler.ServeHTTP(w, req)
		fmt.Println("Executing middleware after response phase")
	})
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing Middleware")
	w.Write([]byte("OK"))
}

func ExecuteMiddleware() {
	// Get Handler Func as closure function
	mainHandler := http.HandlerFunc(handlerFunc)
	http.Handle("/", middleware(mainHandler))
	http.ListenAndServe(":9000", nil)
}
