package routers

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter defines api routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())

	router.GET("/rate", func(c *gin.Context) {
		c.JSON(200, "rate limiting test")
	})

	return router
}
