package middlewares

import (
	"goLandCRUD/logger"
	"goLandCRUD/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		logger.Error("Token not provided")
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}
	logger.Info("Token provided")
	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
