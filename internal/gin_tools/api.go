package gintools

import (
	"errors"

	logmetrics "github.com/CakeForKit/GinFiberBenchmarks.git/internal/log_metrics"
	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/metrics"
	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MetricsRouter struct {
	logger  logmetrics.MetricsLogger
	logsDir string
}

var (
	ErrRequestIDNotFound = errors.New("request_id not found")
	ErrRequestIDNotUUID  = errors.New("request_id is not a uuid")
)

func NewMetricsRouter(router *gin.RouterGroup, logger logmetrics.MetricsLogger, logsDir string) MetricsRouter {
	r := MetricsRouter{
		logger:  logger,
		logsDir: logsDir,
	}
	router.GET("/dump", r.DumpLogs)
	router.GET("/metrics", gin.WrapH(metrics.GetHttpHandler()))

	router.POST("/flat", r.SerializeFlat)
	router.POST("/deep", r.SerializeDeep)
	router.POST("/hierarchy", r.SerializeHierarchy)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})
	return r
}

func (r *MetricsRouter) DumpLogs(c *gin.Context) {
	if err := r.logger.DumpLogs(r.logsDir); err != nil {
		panic(err.Error())
	}
}

func (r *MetricsRouter) SerializeFlat(c *gin.Context) {
	var obj models.FlatStruct
	reqID, ok := c.Get("request_id")
	if !ok {
		c.JSON(500, gin.H{
			"error": ErrRequestIDNotFound.Error(),
		})
	}
	reqIDStr, ok := reqID.(uuid.UUID)
	if !ok {
		c.JSON(500, gin.H{
			"error": ErrRequestIDNotUUID.Error(),
		})
	}
	r.logger.SetSerializeStartTime(reqIDStr)
	err := c.BindJSON(&obj)
	r.logger.SetSerializeEndTime(reqIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func (r *MetricsRouter) SerializeDeep(c *gin.Context) {
	var obj models.TreeNode
	reqID, ok := c.Get("request_id")
	if !ok {
		c.JSON(500, gin.H{
			"error": ErrRequestIDNotFound.Error(),
		})
	}
	reqIDStr, ok := reqID.(uuid.UUID)
	if !ok {
		c.JSON(500, gin.H{
			"error": ErrRequestIDNotUUID.Error(),
		})
	}
	r.logger.SetSerializeStartTime(reqIDStr)
	err := c.BindJSON(&obj)
	r.logger.SetSerializeEndTime(reqIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func (r *MetricsRouter) SerializeHierarchy(c *gin.Context) {
	var obj models.FlatHierarchyStruct
	reqID, ok := c.Get("request_id")
	if !ok {
		c.JSON(500, gin.H{
			"error": ErrRequestIDNotFound.Error(),
		})
	}
	reqIDStr, ok := reqID.(uuid.UUID)
	if !ok {
		c.JSON(500, gin.H{
			"error": ErrRequestIDNotUUID.Error(),
		})
	}
	r.logger.SetSerializeStartTime(reqIDStr)
	err := c.BindJSON(&obj)
	r.logger.SetSerializeEndTime(reqIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
}
