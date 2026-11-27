package router_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/config"
	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/database"
	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/router"
)

func TestHealthReturns503WhenDatabaseDisconnected(t *testing.T) {
	cfg := config.Load()
	h := router.New(cfg, &database.Store{}, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("GET /api/v1/health expected 503 when DB disconnected, got %d", rec.Code)
	}

	var body map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("GET /api/v1/health invalid JSON: %v", err)
	}
	if ok, _ := body["success"].(bool); ok {
		t.Fatalf("expected success=false on 503 response, got %v", body)
	}
}

func TestModuleEndpointsRespond(t *testing.T) {
	cfg := config.Load()
	h := router.New(cfg, &database.Store{}, nil)

	endpoints := []string{
		"/api/v1/core/divisions",
		"/api/v1/dashboard/overview",
		"/api/v1/dashboard/kpi",
		"/api/v1/assets",
		"/api/v1/assets/AST-001/machine-card",
		"/api/v1/work-orders",
		"/api/v1/pm-schedules",
		"/api/v1/inspections",
	}

	for _, ep := range endpoints {
		req := httptest.NewRequest(http.MethodGet, ep, nil)
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("GET %s expected 200 OK, got %d", ep, rec.Code)
		}

		var body map[string]any
		if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
			t.Fatalf("GET %s returned invalid JSON: %v", ep, err)
		}
		if ok, _ := body["success"].(bool); !ok {
			t.Fatalf("GET %s expected success=true, got %v", ep, body)
		}
	}
}

func TestCORSMultiOrigin(t *testing.T) {
	cfg := config.Load()
	cfg.CORSAllowedOrigins = []string{"http://localhost:3000", "http://192.168.10.50:3000"}
	h := router.New(cfg, &database.Store{}, nil)

	req := httptest.NewRequest(http.MethodOptions, "/api/v1/core/divisions", nil)
	req.Header.Set("Origin", "http://192.168.10.50:3000")
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected 204 for CORS preflight, got %d", rec.Code)
	}
	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "http://192.168.10.50:3000" {
		t.Fatalf("expected LAN origin echoed, got %q", got)
	}
}
