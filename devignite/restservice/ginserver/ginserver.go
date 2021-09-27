package ginserver

import (
	"github.com/devignite/mongodb"
	"github.com/gin-gonic/gin"
)

func HostGinServer() {
	r := gin.Default()

	r.GET("/v1/movies/count", func(c *gin.Context) {
		// Gin enables JSON Serializer
		var mongo = mongodb.MongoStruct{}
		c.JSON(200, gin.H{
			"Count": mongo.GetMovieCount(),
		})
	})

	r.GET("/v1/movies/all", func(c *gin.Context) {
		// Gin enables JSON Serializer
		var mongo = mongodb.MongoStruct{}
		c.Writer.Write(mongo.GetMoviesBytes())
	})

	r.Run(":9000")
}
