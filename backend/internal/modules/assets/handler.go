package assets

import (
	"net/http"

	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/database"
	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/shared/response"
)

// Handler owns Track 3 endpoints: Modul 2 Data Aset, Spesifikasi Mesin & Kartu Mesin (URS 3.1–3.6).
type Handler struct {
	store *database.Store
}

func NewHandler(store *database.Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/assets", h.ListAssets)
	mux.HandleFunc("POST /api/v1/assets", h.CreateAsset)
	mux.HandleFunc("GET /api/v1/assets/{id}", h.GetAssetDetail)
	mux.HandleFunc("GET /api/v1/assets/{id}/machine-card", h.GetMachineCard)
	mux.HandleFunc("POST /api/v1/assets/{id}/deactivate", h.RequestDeactivation)
}

func (h *Handler) ListAssets(w http.ResponseWriter, _ *http.Request) {
	response.JSON(w, http.StatusOK, map[string]any{
		"module":  "module-2-assets-machine-card",
		"owner":   "Lovi (Teknik Elektro)",
		"urs_ref": []string{"3.1"},
		"items":   []any{},
		"note":    "Stub endpoint ready for Track 3 (Lovi) PostgreSQL implementation.",
	})
}

func (h *Handler) CreateAsset(w http.ResponseWriter, _ *http.Request) {
	response.Fail(w, http.StatusNotImplemented, "NOT_IMPLEMENTED", "Endpoint pendaftaran aset (URS 3.2) sedang dikerjakan pada branch feat/module-2-assets-machine-card.", nil)
}

func (h *Handler) GetAssetDetail(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]any{
		"asset_id": r.PathValue("id"),
		"urs_ref":  []string{"3.3"},
		"note":     "Detail & Spesifikasi Mesin stub (Track 3 - Lovi).",
	})
}

func (h *Handler) GetMachineCard(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]any{
		"asset_id":   r.PathValue("id"),
		"urs_ref":    []string{"3.4"},
		"history":    []any{},
		"components": []any{},
		"note":       "Kartu Mesin stub (Track 3 - Lovi).",
	})
}

func (h *Handler) RequestDeactivation(w http.ResponseWriter, r *http.Request) {
	response.Fail(w, http.StatusNotImplemented, "NOT_IMPLEMENTED", "Pengajuan penonaktifan aset "+r.PathValue("id")+" (URS 3.6) sedang dikerjakan pada branch feat/module-2-assets-machine-card.", nil)
}
