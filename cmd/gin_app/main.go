package main

import (
	"fmt"

	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/cnfg"
	gintools "github.com/CakeForKit/GinFiberBenchmarks.git/internal/gin_tools"
	logmetrics "github.com/CakeForKit/GinFiberBenchmarks.git/internal/log_metrics"
	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/metrics"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := cnfg.NewDefaultConfig()
	metrics.RegisterMetricCollector(conf.MetricsUpdateTimeMS)
	logger := logmetrics.NewLogger()

	engine := gin.New()
	// engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(gintools.MetricMiddleware(logger, "/metrics", "/dump"))
	apiGroup := engine.Group("/")
	mRouter := gintools.NewMetricsRouter(apiGroup, logger, conf.LogsFilename)
	_ = mRouter

	engine.Run(fmt.Sprintf(":%d", conf.GinPort))
}
