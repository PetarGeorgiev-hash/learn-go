package main

import (
	"net/http"

	"github.com/PetarGeorgiev-hash/learning-go/db"
	"github.com/PetarGeorgiev-hash/learning-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	r := gin.Default()
	r.GET("/events", func(ctx *gin.Context) {
		events := models.GetAllEvents()
		ctx.JSON(http.StatusOK, events)
	})
	r.POST("/events", func(ctx *gin.Context) {
		var event models.Event
		err := ctx.ShouldBindJSON(&event)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		event.Id = 1
		event.UserId = 1
		event.Save()
		ctx.JSON(http.StatusCreated, event)
	})
	r.Run(":8080")
}
