package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jarssin/nps-back/internal/infra/middlewares"
)

func TestCorsMiddleware_OPTIONS(t *testing.T) {
	req := httptest.NewRequest(http.MethodOptions, "/", nil)
	rw := httptest.NewRecorder()
	h := middlewares.CorsMiddleware(func(w http.ResponseWriter, r *http.Request) {})
	h(rw, req)
	if rw.Code != http.StatusNoContent {
		t.Errorf("expected status %d, got %d", http.StatusNoContent, rw.Code)
	}
}
