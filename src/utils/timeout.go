package utils

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func TimeoutNew(h gin.HandlerFunc) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(15*time.Second),
		timeout.WithHandler(h),
		timeout.WithResponse(timeoutResponse),
	)
}

func timeoutResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "timeout")
}
