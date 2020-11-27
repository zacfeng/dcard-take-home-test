package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zacfeng/dcard-take-home-test/utils"
	"gotest.tools/assert"
)

// TestRouter is a gin Engine for API test
var TestRouter *gin.Engine

func init() {
	TestRouter = SetupRouter()
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func setup() {
	client := utils.GetRedisClient()
	iter := client.Scan(0, "*", 0).Iterator()
	for iter.Next() {
		err := client.Del(iter.Val()).Err()
		if err != nil {
			panic(err)
		}
	}
}

func TestRateLimit(t *testing.T) {
	setup()

	oneMinuteAfter := time.Now().Add(time.Duration(1) * time.Minute)

	// pass under 60 requests per minute
	for i := 1; i <= 60; i++ {
		w := performRequest(TestRouter, "GET", "/rate")
		assert.Equal(t, strconv.Itoa(i), w.Body.String())
	}

	for i := 1; i <= 5; i++ {
		w := performRequest(TestRouter, "GET", "/rate")
		assert.Equal(t, "Error", w.Body.String())
	}

	for time.Now().Before(oneMinuteAfter) {
	}

	// pass after the next sleep
	for i := 1; i <= 10; i++ {
		w := performRequest(TestRouter, "GET", "/rate")
		assert.Equal(t, strconv.Itoa(i), w.Body.String())
	}
}
