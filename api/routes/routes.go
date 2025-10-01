package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.GET("/events", getEvents)
	r.POST("/events", createEvent)
	r.GET("events/:id", getSingleEvent)
	r.PUT("events/:id", updateEvent)
	r.DELETE("events/:id", deleteEvent)
}
