package middleware

import (
	"context"
	"net/http"

	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/shared/domain"
	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/shared/response"
)

type contextKey string

const ActorContextKey contextKey = "cmms_actor"

// Actor represents an authenticated user and their authorized scope (URS 1.1 & 1.2).
type Actor struct {
	UserID     string
	Username   string
	FullName   string
	Role       domain.Role
	Divisions  []domain.DivisionCode
	CanApprove bool
	CanVerify  bool
	CanExport  bool
}

func ActorFromContext(ctx context.Context) (Actor, bool) {
	actor, ok := ctx.Value(ActorContextKey).(Actor)
	return actor, ok
}

// RequireRoles enforces role-based access control (Track 1 — Gifari will wire JWT/session verification here).
func RequireRoles(allowed ...domain.Role) func(http.Handler) http.Handler {
	allowedSet := make(map[domain.Role]struct{}, len(allowed))
	for _, role := range allowed {
		allowedSet[role] = struct{}{}
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			actor, ok := ActorFromContext(r.Context())
			if !ok {
				response.Fail(w, http.StatusUnauthorized, "UNAUTHORIZED", "Autentikasi diperlukan untuk mengakses fungsi ini (URS 1.1).", nil)
				return
			}
			if _, ok := allowedSet[actor.Role]; !ok {
				// TODO(Track 1 - Gifari): Catat ACCESS_DENIED ke tabel audit_trails (URS 1.2c)
				response.Fail(w, http.StatusForbidden, "FORBIDDEN", "Akses di luar kewenangan peran atau cakupan divisi Anda (URS 1.2).", nil)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
