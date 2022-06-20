package authentication

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Authenticator() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("ex", "this is a middleware example")
		logrus.Info("Before request")

		c.Next()
		logrus.Info("After request")

		latency := time.Since(t)
		logrus.Info("Middleware latency: ", latency)

		status := c.Writer.Status()
		logrus.Info("Request status after: ", status)
	}
}
