package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zacfeng/dcard-take-home-test/middlewares"
)

// SetupRouter defines api routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middlewares.RateLimit())

	router.GET("/rate", func(c *gin.Context) {
		c.String(http.StatusOK, c.MustGet("rate").(string))
	})

	return router
}

func main() {
	SetupRouter().Run()

}
