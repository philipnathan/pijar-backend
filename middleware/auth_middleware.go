package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get access_token from httponlycookie
		authHeader, err := c.Cookie("access_token")
		fmt.Println(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not found"})
			c.Abort()
			return
		}

		parsedToken, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return secretKey, nil
		})

		if err != nil || !parsedToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// get claims from token
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok || !parsedToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// get user_id from claims
		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}
