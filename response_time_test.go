package responsetime

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
	c.Response().WriteHeader(200)
	assert.Equal(t, "0.000", rec.Header().Get(DefaultResponseTimeConfig.HeaderName))
}
