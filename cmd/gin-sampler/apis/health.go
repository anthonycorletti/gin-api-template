package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

// Healthcheck returns a simple healthcheck json response
func Healthcheck(c *gin.Context) {
	version := os.Getenv("SHORT_SHA")
	if version == "" {
		version = "local"
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "alive and kicking",
		"version": version,
		"time":    time.Now().UTC(),
	})
}
