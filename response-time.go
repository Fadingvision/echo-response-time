package responseTime

import (
	"time"

	"github.com/labstack/echo/v4"
)

type (
	// ResponseTimeConfig defines the config for responseTime middleware.
	ResponseTimeConfig struct {
		Digits int

		HeaderName string

		Suffix bool
	}
)

const (
	// DefaultResponseTimeConfig
	DefaultResponseTimeConfig = ResponseTimeConfig{
		Digits:     3,
		HeaderName: "X-Response-Time",
		Suffix:     false,
	}
)

func ResponseTime() echo.MiddlewareFunc {
	c := DefaultResponseTimeConfig
	return CSRFWithConfig(c)
}

func ResponseTimeWithConfig(config ResponseTimeConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			c.Response().Before(func() {
				elapsed := time.Since(start)
				ms := int64(elapsed / time.Millisecond)
				c.Response().Header().Add("X-Response-Time", ms)
			})

			return next(c)
		}
	}
}
