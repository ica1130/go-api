package middleware

import (
	"net/http"
	"rest-api/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		//AbortWithStatusJSON is a helper function that writes the JSON error message and sets the HTTP status code to 401
		// because this is middleware and we want to stop the request from going to the next handler
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required."})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token."})
		return
	}

	context.Set("userId", userId)

	// by using Next() we tell Gin to continue with the next handler in the chain
	context.Next()
}
