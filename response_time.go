package responsetime

import (
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	// ResponseTimeConfig defines the config for responseTime middleware.
	ResponseTimeConfig struct {
		// Digits ...
		Digits int

		// HeaderName is the header shown in your response
		HeaderName string

		// Suffix indicates whether the time has the `ms` suffix
		Suffix bool
	}
)

var (
	// DefaultResponseTimeConfig defines the default config for responseTime middleware.
	DefaultResponseTimeConfig = ResponseTimeConfig{
		Digits:     3,
		HeaderName: "X-Response-Time",
		Suffix:     false,
	}
)

// ResponseTime is the default middleware
func ResponseTime() echo.MiddlewareFunc {
	c := DefaultResponseTimeConfig
	return ResponseTimeWithConfig(c)
}

// ResponseTimeWithConfig is the  middleware you can custom
func ResponseTimeWithConfig(config ResponseTimeConfig) echo.MiddlewareFunc {
	if config.HeaderName == "" {
		config.HeaderName = DefaultResponseTimeConfig.HeaderName
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			c.Response().Before(func() {
				elapsed := float64(time.Since(start) / time.Millisecond)
				ms := fmt.Sprintf("%."+strconv.Itoa(config.Digits)+"f", elapsed)
				if config.Suffix {
					ms += "ms"
				}
				c.Response().Header().Add(config.HeaderName, ms)
			})

			return next(c)
		}
	}
}
