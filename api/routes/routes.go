package routes

import (
	"github.com/PetarGeorgiev-hash/learning-go/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/events", getEvents)
	r.GET("events/:id", getSingleEvent)

	auth := r.Group("/")
	auth.Use(middlewares.Authenticate)
	auth.POST("/events", createEvent)
	auth.PUT("/events/:id", updateEvent)
	auth.DELETE("events/:id", deleteEvent)
	auth.POST("/events/:id/register", registerForEvent)
	auth.DELETE("/events/:id/register", cancelRegistration)

	r.POST("/signup", createUser)
	r.POST("/login", login)
}
