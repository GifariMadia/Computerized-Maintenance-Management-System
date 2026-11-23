package maintenance

import (
	"net/http"

	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/database"
	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/shared/response"
)

// Handler owns Track 4 endpoints: Modul 3 Breakdown, PM, Inspection, Work Orders & Spare Part (URS 4.1–4.12).
type Handler struct {
	store *database.Store
}

func NewHandler(store *database.Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/work-orders", h.ListWorkOrders)
	mux.HandleFunc("POST /api/v1/breakdowns", h.ReportBreakdown)
	mux.HandleFunc("GET /api/v1/pm-schedules", h.ListPMSchedules)
	mux.HandleFunc("GET /api/v1/inspections", h.ListInspections)
	mux.HandleFunc("POST /api/v1/work-orders/{id}/verify", h.VerifyWorkOrder)
}

func (h *Handler) ListWorkOrders(w http.ResponseWriter, _ *http.Request) {
	response.JSON(w, http.StatusOK, map[string]any{
		"module":  "module-3-maintenance-operations",
		"owner":   "Track 4 (Member #4 / Shared Team)",
		"urs_ref": []string{"4.2", "4.10"},
		"items":   []any{},
		"note":    "Stub endpoint ready for Track 4 PostgreSQL implementation.",
	})
}

func (h *Handler) ReportBreakdown(w http.ResponseWriter, _ *http.Request) {
	response.Fail(w, http.StatusNotImplemented, "NOT_IMPLEMENTED", "Pelaporan breakdown (URS 4.1) sedang dikerjakan pada branch feat/module-3-maintenance-ops.", nil)
}

func (h *Handler) ListPMSchedules(w http.ResponseWriter, _ *http.Request) {
	response.JSON(w, http.StatusOK, map[string]any{
		"urs_ref": []string{"4.4", "4.5"},
		"items":   []any{},
	})
}

func (h *Handler) ListInspections(w http.ResponseWriter, _ *http.Request) {
	response.JSON(w, http.StatusOK, map[string]any{
		"urs_ref": []string{"4.6", "4.7"},
		"items":   []any{},
	})
}

func (h *Handler) VerifyWorkOrder(w http.ResponseWriter, r *http.Request) {
	response.Fail(w, http.StatusNotImplemented, "NOT_IMPLEMENTED", "Verifikasi WO "+r.PathValue("id")+" (URS 4.11–4.12) sedang dikerjakan pada branch feat/module-3-maintenance-ops.", nil)
}
