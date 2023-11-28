package tests

import (
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"restApi/internal/app/config"
	"testing"
)

// TestVersionRoute tests the version route
func TestVersionRoute(t *testing.T) {
	router := config.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/version", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"version\":\"1.0\"}", w.Body.String())
}
