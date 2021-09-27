package main

import (
	"fmt"
	"github.com/devignite/restservice/ginserver"
	"github.com/devignite/restservice/gorestful"
	"github.com/emicklei/go-restful"
	"io"
	"net/http"
	"time"
)

func main() {
	// hostRestful()
	ginserver.HostGinServer()
}

func hostRestful() {
	movieContainer := restful.NewContainer()
	movieContainer.Router(restful.CurlyRouter{})
	gorestful.RestOperations(movieContainer)

	server := http.Server{
		Addr:    ":9000",
		Handler: movieContainer,
	}

	server.ListenAndServe()
}

func createPingService() {
	// Create a web service
	webservice := new(restful.WebService)

	// Create a route and attach it to handler in the service
	webservice.Route(webservice.GET("/ping").To(pingTime))

	// add webservice to the app
	restful.Add(webservice)
	http.ListenAndServe(":9000", nil)
}

func pingTime(r *restful.Request, resp *restful.Response) {
	// Write "demo" to the response
	io.WriteString(resp, fmt.Sprintf("%s", time.Now()))
}
