package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

// StartApp will handle the router and requests to
// this Api
func StartApp() {
	urlMappings()
	router.Run("localhost:8080")
}
