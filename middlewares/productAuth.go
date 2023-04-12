package middlewares

import (
	"challenge-9/models"
	"challenge-9/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProductAuth(productService *services.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		productId, err := strconv.Atoi(c.Param("productId"))

		var mapId map[string]uint

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error_status":  "Bad request",
				"error_message": "Invalid parameter",
			})
			return
		}

		data := c.MustGet("data").(map[string]any)
		userData := data["user"].(models.User)

		product, err := productService.GetUserIdByProductId(uint(productId))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		if product.UserID != userData.ID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		mapId = map[string]uint{
			"productId": uint(productId),
			"userId":    uint(userData.ID),
		}

		c.Set("mapId", mapId)

		c.Next()
	}
}
