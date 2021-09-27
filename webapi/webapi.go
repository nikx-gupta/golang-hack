package main

import (
	"fmt"
	"github.com/devignite/webapi/basicHandler"
	"github.com/devignite/webapi/dataprovider"
	"github.com/julienschmidt/httprouter"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	// Pass handler functions to the path
	http.HandleFunc("/basic", basicHandler.ReportHandler)

	// handle root path and choose dynamic response
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlFragments := strings.Split(r.URL.Path, "/")
		if urlFragments[1] == "report" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlFragments[2]))
			if number >= 1 && number <= 10 {
				fmt.Fprintf(w, "%q", dataprovider.Numerals[number])
			} else {
				w.WriteHeader(http.StatusAccepted)
				w.Write([]byte("Invalid number"))
			}
		}
	})

	// Use Http Router
	router := httprouter.New()
	router.GET("/api/v1/randomInt", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "%f", rand.Float64())
	})
	router.GET("/api/v1/randomInt/:name", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "the param is %s", params.ByName("name"))
	})
	// Declare setting and listen on server
	server := http.Server{Addr: ":8000"}
	log.Fatal(server.ListenAndServe())

	// Uncomment to Use listener directly
	// log.Fatal(http.ListenAndServe(":8000", nil))

	// // Using Mux
	// mux := &simplemux.ManualServerMux{}
	// http.ListenAndServe(":9000", mux)

	// // Using Single Mux handling multiple paths /randomFloat /randomInt
	// http.ListenAndServe(":9000", simplemux.CreateMultiMux())

	// // Static File Server
	// router.ServeFiles("/static/*filepath", http.Dir("/home/nikx/src/lab-golang/sample-static"))

	// // Use Router
	// http.ListenAndServe(":9000", router)

	// // Gorilla Mux
	// gorillamux.RunGorillaMux()

	// // Chain Middleware
	// middleware.ExecuteChainMiddleware()

	// // Logging Middleware
	// middleware.RunWithLoggingMiddleware()
}
