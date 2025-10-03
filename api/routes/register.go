package routes

import (
	"net/http"
	"strconv"

	"github.com/PetarGeorgiev-hash/learning-go/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(ctx *gin.Context) {
	currentUserId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	err = event.Register(currentUserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Registered"})
}

func cancelRegistration(ctx *gin.Context) {
	currentUserId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	err = event.Cancel(currentUserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Deleted"})
}
