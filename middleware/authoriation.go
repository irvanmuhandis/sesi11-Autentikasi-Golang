package middleware

import (
	"net/http"
	"sesi11/database"
	"sesi11/model"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		productId, err := strconv.Atoi(c.Param("productId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":    "Bad Request",
				"messsage": "Invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		product := model.Product{}

		err = db.Select("user_id").First(&product, uint(productId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":    "Not Found",
				"messsage": "Product not found",
			})
			return
		}

		if product.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":    "Unauthorized",
				"messsage": "You didnt have acces",
			})
			return
		}

		// Lanjut ke nedpoint berikutnya
		c.Next()
	}
}
