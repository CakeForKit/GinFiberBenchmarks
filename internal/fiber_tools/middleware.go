package fibertools

import (
	"slices"

	logmetrics "github.com/CakeForKit/GinFiberBenchmarks.git/internal/log_metrics"
	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/metrics"
	"github.com/gofiber/fiber/v2"
)

func MetricMiddleware(logger logmetrics.MetricsLogger, exceptPaths ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentPath := c.Path()
		if slices.Contains(exceptPaths, currentPath) {
			return c.Next()
		}
		requestID := logger.CreateRequest()
		logger.SetRequestPath(requestID, currentPath)
		c.Locals("request_id", requestID)

		metrics.TotalHttpRequestCounter.WithLabelValues(c.Method(), currentPath).Inc()
		return c.Next()
	}
}
