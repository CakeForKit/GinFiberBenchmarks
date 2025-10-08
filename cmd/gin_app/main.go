package main

import (
	"fmt"
	"log"
	"time"

	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/cnfg"
	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/models"
	"github.com/gin-gonic/gin"
)

func main() {
	config := cnfg.NewDefaultConfig()
	simple, tree, wide := models.LoadTestData()

	r := gin.Default()

	// Middleware для метрик
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()

		fmt.Printf("METRIC: path=%s method=%s duration=%f\n",
			c.Request.URL.Path, c.Request.Method, duration)
	})

	r.GET("/simple", func(c *gin.Context) {
		c.JSON(200, simple)
	})

	r.GET("/tree", func(c *gin.Context) {
		c.JSON(200, tree)
	})

	r.GET("/wide", func(c *gin.Context) {
		c.JSON(200, wide)
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})
	log.Printf("Gin server starting on http://%s:%d", config.Host, config.GinPort)
	r.Run(fmt.Sprintf(":%d", config.GinPort))
}
