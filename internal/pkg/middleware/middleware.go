package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/FauzanAr/fleet-management/internal/pkg/logger"
)

func GinRequestTrace(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		traceID := uuid.New().String()
		ctx := context.WithValue(c.Request.Context(), "trace_id", traceID)
		ctx = context.WithValue(ctx, "url", c.Request.RequestURI)
		ctx = context.WithValue(ctx, "method", c.Request.Method)
		ctx = context.WithValue(ctx, "remote_ip", c.ClientIP())
		c.Request = c.Request.WithContext(ctx)

		log.Info(ctx, "Request started", nil)
		c.Header("X-Trace-ID", traceID)

		c.Next()

		log.Info(ctx, "Request completed", logger.MetaData{
			"status":       c.Writer.Status(),
			"elapsed_time": time.Since(start).Seconds(),
		})
	}
}
