package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/fcc-timestamp-ms/handlers"
	"github.com/madjiebimaa/fcc-timestamp-ms/middlewares"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middlewares.CORS())

	r.GET("/api/:unix", handlers.UnixToUTCHandler)
	if err := r.Run(":3000"); err != nil {
		log.Fatal("can't connect to server at port 3000")
	}
}
