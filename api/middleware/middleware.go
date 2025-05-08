package middleware

import (
	"net/http"
	"strings"
	"todo/api/utils"
	"todo/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Authorization header missing"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenStr, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return utils.JwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid token"})
			c.Abort()
			return
		}

		claims := token.Claims.(*utils.Claims)
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}

// CORSMiddleware handles cross-origin requests
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Asosiy CORS headerlar
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // yoki domen: http://localhost:3000
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Preflight so‘rovga javob
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
