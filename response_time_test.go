package responsetime

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestResponseTime(t *testing.T) {
	e := echo.New()

	// Wildcard origin
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := ResponseTime()(echo.NotFoundHandler)
	h(c)
	time.Sleep(time.Second)
	c.Response().WriteHeader(200)
	responseTime, _ := strconv.ParseFloat(rec.Header().Get(DefaultResponseTimeConfig.HeaderName), 64)
	assert.True(t, responseTime >= 1000)
}
