package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/vearne/cache-test/resource"
	"strconv"
	"strings"
	"time"
)

const (
	PromLabelMethod = "method" //http method
	PromLabelPath   = "path"   //http path
	PromLabelStatus = "status" //http response status
)

// Metric metric middleware
func Metric() gin.HandlerFunc {
	return func(c *gin.Context) {
		tBegin := time.Now()
		c.Next()

		duration := float64(time.Since(tBegin)) / float64(time.Second)

		path := parsePath(c.Request.URL.Path)
		resource.HTTPReqTotal.With(prometheus.Labels{
			PromLabelMethod: c.Request.Method,
			PromLabelPath:   path,
			PromLabelStatus: strconv.Itoa(c.Writer.Status()),
		}).Inc()

		resource.HTTPReqDuration.With(prometheus.Labels{
			PromLabelMethod: c.Request.Method,
			PromLabelPath:   path,
		}).Observe(duration)

	}
}

func parsePath(path string) string {
	itemList := strings.Split(path, "/")
	return strings.Join(itemList[0:4], "/")
}
