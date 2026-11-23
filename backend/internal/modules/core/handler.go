package core

import (
	"net/http"

	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/database"
	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/shared/response"
)

// Handler owns Track 1 endpoints: Health, Auth (URS 1.1), RBAC/Scopes (URS 1.2), Audit Trail (URS 1.4).
type Handler struct {
	store *database.Store
}

func NewHandler(store *database.Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/health", h.Health)
	mux.HandleFunc("GET /api/v1/core/divisions", h.ListDivisions)
	mux.HandleFunc("GET /api/v1/core/audit-trails", h.ListAuditTrails)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	dbStatus := "disconnected"
	if h.store.Healthy(r.Context()) {
		dbStatus = "connected"
	}
	response.JSON(w, http.StatusOK, map[string]any{
		"service":  "cmms-backend",
		"site":     "PT Bintang Toedjoe Site Pulogadung",
		"status":   "ok",
		"database": dbStatus,
		"timezone": "Asia/Jakarta (WIB)",
	})
}

func (h *Handler) ListDivisions(w http.ResponseWriter, _ *http.Request) {
	// Baseline master response matching Migration 001 (Track 1 wires live PostgreSQL query).
	response.JSON(w, http.StatusOK, []map[string]string{
		{"code": "MAINTENANCE", "name": "Maintenance"},
		{"code": "UTILITY", "name": "Utility"},
		{"code": "BUILDING", "name": "Building Maintenance"},
		{"code": "SPARE_PART", "name": "Spare Part"},
	})
}

func (h *Handler) ListAuditTrails(w http.ResponseWriter, _ *http.Request) {
	response.JSON(w, http.StatusOK, map[string]any{
		"module":  "core-audit-trail",
		"urs_ref": []string{"1.4", "ALCOA-1.1", "ALCOA-1.4"},
		"items":   []any{},
		"note":    "Stub endpoint ready for Track 1 (Gifari) PostgreSQL implementation.",
	})
}
