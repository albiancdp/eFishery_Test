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

	v1User := r.Group("/api/v1/user")
	v1Admin := r.Group("/api/v1/admin")
	v1User.Use(JwtMiddleware.UserRole).GET("resource/fetch", ResourceController.FetchResource)
	v1Admin.Use(JwtMiddleware.AdminRole).GET("resource/aggregate", ResourceController.FetchResource)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.Run(":3001")
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
