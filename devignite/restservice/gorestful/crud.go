package gorestful

import (
	"github.com/devignite/mongodb"
	"github.com/emicklei/go-restful"
)

func RestOperations(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/v1/movies").
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/count").To(mongodb.GetMovieCount))

	container.Add(ws)
}

