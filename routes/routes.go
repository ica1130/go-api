package routes

import (
	"rest-api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	// we can use the middleware to protect the route by placing the Auth fucntion before the handler function
	// authenticated.POST("/events", middleware.Authenticate, createEvent)

	// we can also protect a group of routes by using the Group method
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
