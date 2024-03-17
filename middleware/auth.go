package middleware

import (
	"net/http"
	"sesi11/helpers"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":    "Unauthorized",
				"messsage": err.Error(),
			})
			return
		}
		//simpan token ke data request untuk diambil endpoint slanjutnya
		c.Set("userData", verifyToken)
		// Lanjut ke nedpoint berikutnya
		c.Next()
	}
}
