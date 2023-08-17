package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.POST("/webhook/ec-notification", func(c *gin.Context) {

		var notification map[string]interface{}
		if decodeErr := json.NewDecoder(c.Request.Body).Decode(&notification); decodeErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid request body",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"notification-data": notification,
		})
	})
	_ = r.Run()
}
