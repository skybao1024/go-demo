package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"myproject/internal/config"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "Authorization header is required"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		claims, err := validateToken(parts[1], cfg.JWTSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

func validateToken(tokenString string, jwtSecret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func GetUserIDFromContext(c *gin.Context) (uint, bool) {
	claims, exists := c.Get("claims")
	if !exists {
		return 0, false
	}

	mapClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		return 0, false
	}

	userID, ok := mapClaims["user_id"].(float64)
	if !ok {
		return 0, false
	}

	return uint(userID), true
}
