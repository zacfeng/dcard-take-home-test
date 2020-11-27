package middlewares

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zacfeng/dcard-take-home-test/utils"
)

// RateLimit limit 60 rpm request
func RateLimit() gin.HandlerFunc {

	client := utils.GetRedisClient()

	return func(c *gin.Context) {
		key := c.ClientIP()

		cnt, err := client.Incr(key).Result()
		if err == nil && cnt == 1 { // the key was just created
			err = client.ExpireAt(key, time.Now().Add(1*time.Minute)).Err()
		}

		// reach reate limit
		if cnt > 60 {
			c.String(http.StatusOK, "Error")
			c.Abort()
			return
		}

		if err != nil {
			c.AbortWithStatus(http.StatusServiceUnavailable)
		}

		ttl, ttlErr := client.TTL(key).Result()

		if ttlErr != nil {
			c.AbortWithStatus(http.StatusServiceUnavailable)
		}

		c.Header("X-RateLimit-Limit", strconv.FormatInt(60, 10))
		c.Header("X-RateLimit-Remaining", strconv.FormatInt(60-cnt, 10))
		c.Header("X-RateLimit-Reset", fmt.Sprintf("%.0f", ttl.Seconds()))

		c.Set("rate", strconv.FormatInt(cnt, 10))

		c.Next()
	}
}
