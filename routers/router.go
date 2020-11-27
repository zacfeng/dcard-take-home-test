package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter defines api routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())

	router.GET("/rate", func(c *gin.Context) {
		c.String(http.StatusOK, "rate limiting test")
	})

	return router
}
