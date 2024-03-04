package helpers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/daniadjt/task-5-pbi-btpns-dania-djatniko/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	expirationTime := int64(claims["exp"].(float64))
	if time.Now().Unix() > expirationTime {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userID, ok := claims["sub"].(float64)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User
	err = models.DB.Where("user_id = ?", int(userID)).First(&user).Error
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("user_id", user.UserID)
	c.Next()
}
