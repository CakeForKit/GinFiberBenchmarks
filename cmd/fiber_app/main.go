package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/cnfg"
	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := cnfg.NewDefaultConfig()
	simple, tree, wide := models.LoadTestData()

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Middleware для метрик
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start).Seconds()

		fmt.Printf("METRIC: path=%s method=%s status=%d duration=%f\n",
			c.Path(), c.Method(), c.Response().StatusCode(), duration)
		return err
	})

	app.Get("/simple", func(c *fiber.Ctx) error {
		return c.JSON(simple)
	})

	app.Get("/tree", func(c *fiber.Ctx) error {
		return c.JSON(tree)
	})

	app.Get("/wide", func(c *fiber.Ctx) error {
		return c.JSON(wide)
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "healthy"})
	})

	log.Printf("Fiber server starting on http://%s:%d", config.Host, config.FiberPort)
	app.Listen(fmt.Sprintf(":%d", config.FiberPort))
}
