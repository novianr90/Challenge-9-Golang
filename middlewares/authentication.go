package middlewares

import (
	"challenge-9/helpers"
	"challenge-9/services"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication(us *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)

		superUser := os.Getenv("EMAIL_ADMIN")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error_status":  "Unauthenticated",
				"error_message": err.Error(),
			})
			return
		}

		email := verifyToken.(jwt.MapClaims)["email"].(string)

		user, err := us.GetUserByEmail(email)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		mapData := map[string]any{
			"isAdmin": user.Email == superUser,
			"user":    user,
		}

		c.Set("data", mapData)

		c.Next()
	}
}
