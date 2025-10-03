package middlewares

import (
	"net/http"

	"github.com/PetarGeorgiev-hash/learning-go/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No token in header"})
		return
	}
	userId, err := utils.VerifyJwt(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}
	ctx.Set("userId", userId)
	ctx.Next()
}
