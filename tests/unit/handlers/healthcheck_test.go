package handlers

import (
	"github.com/stretchr/testify/assert"
	"github.com/zett-8/go-clean-echo/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	e := handlers.Echo()

	req := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, handlers.HealthCheckHandler(c)) {
		assert.Equal(t, rec.Code, http.StatusOK)
		assert.Equal(t, rec.Body.String(), "OK")
	}
}
