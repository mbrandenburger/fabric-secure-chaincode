package main

import (
	"flag"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	startServer()
}

var flagPort string

func init() {
	flag.StringVar(&flagPort, "port", "3000", "Port to listen on")
}

func startServer() {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "x-user")

	r := gin.Default()
	r.Use(cors.New(config))

	// controller
	r.POST("/api/approve", approve)

	r.Run(":" + flagPort)
}

func approve(c *gin.Context) {

	// TODO

	// return answer
	c.IndentedJSON(http.StatusOK, "approved")
}
