package middleware

import (
	"fmt"
	"log"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Executing Main Handler")
}

func verifyHeaderToken(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate Heaader Token and short circuit the pipeline if not present
		if r.Header.Get("x-token") == "" {
			w.Write([]byte("Token Not Present"))
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func addTraceId(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add TraceId to the Response Header
		w.Header().Add("traceId", "testId")
		handler.ServeHTTP(w, r)
	})
}

func ExecuteChainMiddleware() {
	handler := http.HandlerFunc(mainHandler)
	log.Fatal(http.ListenAndServe(":9000", verifyHeaderToken(addTraceId(handler))))
}
