package routes

import (
	"net/http"

	"github.com/PetarGeorgiev-hash/learning-go/models"
	"github.com/PetarGeorgiev-hash/learning-go/utils"
	"github.com/gin-gonic/gin"
)

func createUser(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	token, err := utils.GenerateJwt(user.Email, user.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, token)
}
