package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping : HeartBeat function to check the server status
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}
