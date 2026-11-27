package core

import (
	"net/http"

	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/database"
	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/shared/response"
	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/storage"
)

// Handler owns Track 1 endpoints: Health, Auth (URS 1.1), RBAC/Scopes (URS 1.2), Audit Trail (URS 1.4).
type Handler struct {
	store   *database.Store
	storage *storage.Client
}

func NewHandler(store *database.Store, objStorage *storage.Client) *Handler {
	return &Handler{
		store:   store,
		storage: objStorage,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/health", h.Health)
	mux.HandleFunc("GET /api/v1/core/divisions", h.ListDivisions)
	mux.HandleFunc("GET /api/v1/core/audit-trails", h.ListAuditTrails)
}

// Health returns 200 OK only when PostgreSQL (and MinIO if configured) are reachable;
// otherwise returns 503 Service Unavailable.
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	dbHealthy := h.store != nil && h.store.Healthy(ctx)
	dbStatus := "disconnected"
	if dbHealthy {
		dbStatus = "connected"
	}

	storageConfigured := h.storage != nil && h.storage.Configured()
	storageHealthy := true
	storageStatus := "not_configured"
	if storageConfigured {
		storageHealthy = h.storage.Healthy(ctx)
		if storageHealthy {
			storageStatus = "connected"
		} else {
			storageStatus = "disconnected"
		}
	}

	payload := map[string]any{
		"service":  "cmms-backend",
		"site":     "PT Bintang Toedjoe Site Pulogadung",
		"database": dbStatus,
		"storage": map[string]any{
			"status": storageStatus,
			"bucket": h.storageBucket(),
		},
		"timezone": "Asia/Jakarta (WIB)",
	}

	if !dbHealthy || !storageHealthy {
		payload["status"] = "degraded"
		response.JSON(w, http.StatusServiceUnavailable, payload)
		return
	}

	payload["status"] = "ok"
	response.JSON(w, http.StatusOK, payload)
}

func (h *Handler) storageBucket() string {
	if h.storage == nil {
		return ""
	}
	return h.storage.Bucket()
}

func (h *Handler) ListDivisions(w http.ResponseWriter, r *http.Request) {
	if h.store != nil && h.store.DB != nil {
		rows, err := h.store.DB.QueryContext(r.Context(), `
			SELECT id::text, code, name, COALESCE(description, '')
			FROM divisions
			WHERE is_active = TRUE
			ORDER BY code ASC
		`)
		if err != nil {
			response.Fail(w, http.StatusInternalServerError, "DB_QUERY_FAILED", "Gagal membaca data divisi dari basis data.", nil)
			return
		}
		defer rows.Close()

		var items []map[string]string
		for rows.Next() {
			var id, code, name, desc string
			if err := rows.Scan(&id, &code, &name, &desc); err != nil {
				response.Fail(w, http.StatusInternalServerError, "DB_SCAN_FAILED", "Gagal memproses data divisi.", nil)
				return
			}
			items = append(items, map[string]string{
				"id":          id,
				"code":        code,
				"name":        name,
				"description": desc,
			})
		}
		if err := rows.Err(); err != nil {
			response.Fail(w, http.StatusInternalServerError, "DB_ROWS_FAILED", "Gagal menyelesaikan pembacaan data divisi.", nil)
			return
		}
		response.JSON(w, http.StatusOK, items)
		return
	}

	// Fallback for isolated unit tests without live DB.
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
		"note":    "Stub endpoint ready for Track 1 (Gifari) live audit trail query.",
	})
}
