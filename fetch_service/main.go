package main

import (
	"fetch_service/controllers"
	"fetch_service/middleware"

	"github.com/gin-gonic/gin"
)

var ResourceController = new(controllers.ResourceController)
var JwtMiddleware = new(middleware.JwtMiddleware)

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	v1 := r.Group("/api/v1")
	v1.Use(JwtMiddleware.JwtTokenCheck)
	v1.GET("resource/fetch", ResourceController.FetchResource)

	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
