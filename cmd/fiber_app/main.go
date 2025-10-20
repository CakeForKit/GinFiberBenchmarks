package main

import (
	"encoding/json"
	"fmt"

	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/cnfg"
	fibertools "github.com/CakeForKit/GinFiberBenchmarks.git/internal/fiber_tools"
	logmetrics "github.com/CakeForKit/GinFiberBenchmarks.git/internal/log_metrics"
	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/metrics"
	"github.com/gofiber/fiber/v2"
)

func main() {
	conf := cnfg.NewDefaultConfig()
	metrics.RegisterMetricCollector(conf.MetricsUpdateTimeMS)
	logger := logmetrics.NewLogger()

	engine := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	engine.Use(fibertools.MetricMiddleware(logger, "/metrics", "/dump"))
	apiGroup := engine.Group("/")
	mRouter := fibertools.NewMetricsRouter(apiGroup, logger, conf.LogsFilename)
	_ = mRouter

	engine.Listen(fmt.Sprintf(":%d", conf.FiberPort))
}

/*
введение
моделирование
керхова

симметричные и блочные алгоритмы
поточное шифрование
энигма
открытый ключа
фон неймана

*/
