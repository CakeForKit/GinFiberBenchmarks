package gintools

import (
	"slices"

	logmetrics "github.com/CakeForKit/GinFiberBenchmarks.git/internal/log_metrics"
	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/metrics"
	"github.com/gin-gonic/gin"
)

func MetricMiddleware(logger logmetrics.MetricsLogger, exceptPaths ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if slices.Contains(exceptPaths, c.Request.URL.Path) {
			c.Next()
			return
		}
		requestID := logger.CreateRequest()
		logger.SetRequestPath(requestID, c.Request.URL.Path)
		c.Set("request_id", requestID)

		// metrics.ActiveHttpRequestCounter.WithLabelValues(c.Request.Method, c.Request.URL.Path).Inc()
		metrics.TotalHttpRequestCounter.WithLabelValues(c.Request.Method, c.Request.URL.Path).Inc()

		// defer func() {
		// 	// logger.SetResponseStatus(requestID, c.Writer.Status())
		// 	metrics.ActiveHttpRequestCounter.WithLabelValues(c.Request.Method, c.Request.URL.Path).Dec()
		// }()

		c.Next()
	}
}
