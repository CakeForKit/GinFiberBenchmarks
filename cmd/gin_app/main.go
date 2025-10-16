package main

import (
	"fmt"
	"os"
	"path/filepath"

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

	currentDir, err := os.Getwd()
	if err != nil {
		panic("Failed to get current directory: " + err.Error())
	}
	logsPath := filepath.Join(currentDir, conf.LogsFilename)
	fmt.Printf("logsPath: %s\n\n", logsPath)

	f, err := os.Create(logsPath)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(gintools.MetricMiddleware(logger, "/metrics", "/dump"))
	apiGroup := engine.Group("/")
	mRouter := gintools.NewMetricsRouter(apiGroup, logger, f)
	_ = mRouter

	engine.Run(fmt.Sprintf(":%d", conf.GinPort))
}

/*

var (
	ginRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "gin_request_duration_seconds",
			Help:    "Duration of HTTP requests in Gin",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"handler", "method"},
	)

	ginRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gin_requests_total",
			Help: "Total number of requests in Gin",
		},
		[]string{"handler", "method", "status"},
	)
)

func init() {
	prometheus.MustRegister(ginRequestDuration, ginRequestsTotal)
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	r := gin.New()
	r.Use(gin.Recovery())

	// Metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Test endpoints
	r.GET("/user", ginWrapMetrics(getUser, "get_user"))
	r.POST("/user", ginWrapMetrics(createUser, "create_user"))

	r.Run(":8080")
}

func getUser(c *gin.Context) {
	user := User{
		ID:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		Active:    true,
		CreatedAt: time.Now(),
	}
	c.JSON(200, user)
}

func createUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, user)
}

func ginWrapMetrics(handler gin.HandlerFunc, handlerName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		handler(c)

		duration := time.Since(start).Seconds()
		status := c.Writer.Status()

		ginRequestDuration.WithLabelValues(handlerName, c.Request.Method).Observe(duration)
		ginRequestsTotal.WithLabelValues(handlerName, c.Request.Method, fmt.Sprintf("%d", status)).Inc()
	}
}

*/

// func main() {
// 	config := cnfg.NewDefaultConfig()
// 	simple, tree, wide := models.LoadTestData()

// 	r := gin.Default()

// 	// Middleware для метрик
// 	r.Use(func(c *gin.Context) {
// 		start := time.Now()
// 		c.Next()
// 		duration := time.Since(start).Seconds()

// 		fmt.Printf("METRIC: path=%s method=%s duration=%f\n",
// 			c.Request.URL.Path, c.Request.Method, duration)
// 	})

// 	r.GET("/simple", func(c *gin.Context) {
// 		c.JSON(200, simple)
// 	})

// 	r.GET("/tree", func(c *gin.Context) {
// 		c.JSON(200, tree)
// 	})

// 	r.GET("/wide", func(c *gin.Context) {
// 		c.JSON(200, wide)
// 	})

// 	r.GET("/health", func(c *gin.Context) {
// 		c.JSON(200, gin.H{"status": "healthy"})
// 	})
// 	log.Printf("Gin server starting on http://%s:%d", config.Host, config.GinPort)
// 	r.Run(fmt.Sprintf(":%d", config.GinPort))
// }
