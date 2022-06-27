package middlewares

import (
	"logging"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var DebugLogger logging.DebugLogger

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		jwt_token, error := context.Request.Cookie("jwt-token")
		if jwt_token == nil || error != nil {
			DebugLogger.Println("User Is Not Authenticated.")
			return
		} else {
			if ok := CheckJWTValidToken(jwt_token.Value); ok {
				context.Next()
			}
		}
	}
}

func CheckJWTValidToken(tokenString string) bool {
	jwt_token := jwt.Parse(tokenString)
	if tokenPayload, ok := jwt_token.Claims.(jwt.tokenString); ok {
		return true
	} else {
		return false
	}
}
