package middlewares 

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"dblab/questlist/models"
	"github.com/golang-jwt/jwt/v4"
	"dblab/questlist/initializers"
	"os"
	"time"
)

func RequireAuth(c *gin.Context) {
	// Get the token from cookie and check if it's valid
	token, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You must be logged in!"})
		c.Abort()
		return
	}

	// Check if the token is valid
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You must be logged in!"})
		c.Abort()
		return 

		// Check if expired token
		if claims["exp"].(float64) < float64(time.Now().Unix()) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Your token has expired!"})
			c.Abort()
			return
		}

		// Check if the user exists
		var user models.User
		if err := initializers.DB.Where("id = ?", claims["id"]).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You must be logged in!"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}


func RequireAdmin(c *gin.Context) {
	// Get the token from cookie and check if it's valid
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You must be logged in!"})
		c.Abort()
		return
	}

	// Check if the token is valid
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You must be logged in!"})
		c.Abort()
		return 

		// Check if expired token
		if claims["exp"].(float64) < float64(time.Now().Unix()) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Your token has expired!"})
			c.Abort()
			return
		}

		// Check if the user exists
		var user models.User
		if err := initializers.DB.Where("id = ?", claims["id"]).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You must be logged in!"})
			c.Abort()
			return
		}
		if user.Role == "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You must be an admin!"})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}

}
// 
