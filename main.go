package main

import (
	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/fcc-timestamp-ms/handlers"
	"github.com/madjiebimaa/fcc-timestamp-ms/middlewares"
)

func main() {
	r := gin.New()
	r.Use(middlewares.CORS())
	r.GET("/api/:unix", handlers.UnixToUTCHandler)
	r.Run(":3000")
}
