package dashboard

import (
	"net/http"

	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/database"
	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/shared/response"
)

// Handler owns Track 2 endpoints: Modul 1 Dashboard & Analitik (URS 2.1–2.9) + Pelaporan (URS 5.1–5.5).
type Handler struct {
	store *database.Store
}

func NewHandler(store *database.Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/dashboard/overview", h.Overview)
	mux.HandleFunc("GET /api/v1/dashboard/kpi", h.KPI)
	mux.HandleFunc("GET /api/v1/dashboard/breakdown-downtime", h.BreakdownDowntime)
	mux.HandleFunc("GET /api/v1/dashboard/costs", h.MaintenanceCosts)
	mux.HandleFunc("GET /api/v1/dashboard/spare-parts", h.SparePartAnalysis)
}

func (h *Handler) Overview(w http.ResponseWriter, _ *http.Request) {
	response.JSON(w, http.StatusOK, map[string]any{
		"module":  "module-1-dashboard-analytics",
		"owner":   "Rainer (Teknik Industri)",
		"urs_ref": []string{"2.1", "2.2", "2.3"},
		"summary": map[string]any{
			"total_assets":        nil, // ALCOA 1.2: null dibedakan dari nilai 0 yang sah
			"open_breakdowns":     nil,
			"overdue_work_orders": nil,
			"pm_compliance_pct":   nil,
		},
		"note": "Stub endpoint ready for Track 2 (Rainer) analytical queries.",
	})
}

func (h *Handler) KPI(w http.ResponseWriter, _ *http.Request) {
	response.JSON(w, http.StatusOK, map[string]any{
		"urs_ref": []string{"2.4", "2.7"},
		"metrics": []map[string]any{
			{"code": "MTTR", "value": nil, "unit": "hours", "status": "DATA_UNAVAILABLE"},
			{"code": "MTBF", "value": nil, "unit": "hours", "status": "DATA_UNAVAILABLE"},
			{"code": "PM_COMPLIANCE", "value": nil, "unit": "%", "status": "DATA_UNAVAILABLE"},
		},
	})
}

func (h *Handler) BreakdownDowntime(w http.ResponseWriter, _ *http.Request) {
	response.JSON(w, http.StatusOK, map[string]any{
		"urs_ref": []string{"2.4"},
		"items":   []any{},
	})
}

func (h *Handler) MaintenanceCosts(w http.ResponseWriter, _ *http.Request) {
	response.JSON(w, http.StatusOK, map[string]any{
		"urs_ref": []string{"2.5", "2.9"},
		"items":   []any{},
	})
}

func (h *Handler) SparePartAnalysis(w http.ResponseWriter, _ *http.Request) {
	response.JSON(w, http.StatusOK, map[string]any{
		"urs_ref": []string{"2.6"},
		"items":   []any{},
	})
}
