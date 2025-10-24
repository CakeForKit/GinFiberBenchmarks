package fibertools

import (
	"errors"
	"fmt"

	logmetrics "github.com/CakeForKit/GinFiberBenchmarks.git/internal/log_metrics"
	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/metrics"
	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/google/uuid"
)

type MetricsRouter struct {
	logger       logmetrics.MetricsLogger
	logsFilename string
}

var (
	ErrRequestIDNotFound = errors.New("request_id not found")
	ErrRequestIDNotUUID  = errors.New("request_id is not a uuid")
)

func NewMetricsRouter(router fiber.Router, logger logmetrics.MetricsLogger, logsFilename string) MetricsRouter {
	r := MetricsRouter{
		logger:       logger,
		logsFilename: logsFilename,
	}

	router.Get("/dump", r.DumpLogs)
	router.Get("/metrics", adaptor.HTTPHandler(metrics.GetHttpHandler()))

	router.Post("/flat", r.SerializeFlat)
	router.Post("/deep", r.SerializeDeep)
	router.Post("/hierarchy", r.SerializeHierarchy)

	router.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "healthy"})
	})

	return r
}

func (r *MetricsRouter) DumpLogs(c *fiber.Ctx) error {
	if err := r.logger.DumpLogs(r.logsFilename); err != nil {
		fmt.Printf("Error: %v\n", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func (r *MetricsRouter) SerializeFlat(c *fiber.Ctx) error {
	var obj models.FlatStruct

	// Получение request_id из контекста
	reqID := c.Locals("request_id")
	if reqID == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": ErrRequestIDNotFound.Error(),
		})
	}

	reqIDStr, ok := reqID.(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": ErrRequestIDNotUUID.Error(),
		})
	}

	r.logger.SetSerializeStartTime(reqIDStr)

	if err := c.BodyParser(&obj); err != nil {
		r.logger.SetSerializeEndTime(reqIDStr)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	r.logger.SetSerializeEndTime(reqIDStr)
	return c.SendStatus(fiber.StatusOK)
}

func (r *MetricsRouter) SerializeDeep(c *fiber.Ctx) error {
	var obj models.TreeNode

	reqID := c.Locals("request_id")
	if reqID == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": ErrRequestIDNotFound.Error(),
		})
	}

	reqIDStr, ok := reqID.(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": ErrRequestIDNotUUID.Error(),
		})
	}

	r.logger.SetSerializeStartTime(reqIDStr)

	if err := c.BodyParser(&obj); err != nil {
		r.logger.SetSerializeEndTime(reqIDStr)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	r.logger.SetSerializeEndTime(reqIDStr)
	return c.SendStatus(fiber.StatusOK)
}

func (r *MetricsRouter) SerializeHierarchy(c *fiber.Ctx) error {
	var obj models.FlatHierarchyStruct

	reqID := c.Locals("request_id")
	if reqID == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": ErrRequestIDNotFound.Error(),
		})
	}

	reqIDStr, ok := reqID.(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": ErrRequestIDNotUUID.Error(),
		})
	}

	r.logger.SetSerializeStartTime(reqIDStr)

	if err := c.BodyParser(&obj); err != nil {
		r.logger.SetSerializeEndTime(reqIDStr)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	r.logger.SetSerializeEndTime(reqIDStr)
	return c.SendStatus(fiber.StatusOK)
}
