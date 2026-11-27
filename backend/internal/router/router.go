package router

import (
	"net/http"

	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/config"
	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/database"
	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/middleware"
	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/modules/assets"
	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/modules/core"
	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/modules/dashboard"
	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/modules/maintenance"
	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/storage"
)

func New(cfg config.Config, store *database.Store, objStorage *storage.Client) http.Handler {
	mux := http.NewServeMux()

	// Track 1 (Gifari — IT): Core Auth, RBAC, Audit Trail, Health, Object Storage
	core.NewHandler(store, objStorage).RegisterRoutes(mux)

	// Track 2 (Rainer — Teknik Industri): Modul 1 Dashboard, KPI, Cost & Reporting
	dashboard.NewHandler(store).RegisterRoutes(mux)

	// Track 3 (Lovi — Teknik Elektro): Modul 2 Master Data Aset, Spesifikasi Mesin & Kartu Mesin
	assets.NewHandler(store).RegisterRoutes(mux)

	// Track 4 (Member #4 / Shared): Modul 3 Breakdown, PM, Inspection & Spare Part
	maintenance.NewHandler(store).RegisterRoutes(mux)

	handler := middleware.AuditRequestLogger(mux)
	handler = middleware.CORS(cfg.CORSAllowedOrigins, handler)
	return handler
}
