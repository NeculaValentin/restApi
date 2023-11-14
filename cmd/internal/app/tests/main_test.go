package tests

import (
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"restApi/cmd/internal/app/config"
	"testing"
)

// TestPingRoute tests the ping route
func TestPingRoute(t *testing.T) {
	router := config.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/version", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
