package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ValidateApiKey(keyName string) gin.HandlerFunc {
	return func(c *gin.Context) {

		var apiKey string = c.GetHeader("x-api-key")
		/*CHECK API KEY*/
		if apiKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "missing credentials",
			})
			return
		}

		if apiKey != os.Getenv(keyName) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid credentials",
			})
			return
		}

		fmt.Println("[ValidateApiKey] OK")
		c.Next()
	}
}
