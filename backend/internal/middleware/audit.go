package middleware

import (
	"log"
	"net/http"
	"time"
)

// AuditRequestLogger records request latency and mutation metadata.
// Track 1 (Gifari) extends this with persistent writes to `audit_trails` for GxP/ALCOA (URS 1.4).
func AuditRequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[CMMS-API] %s %s (%s) idempotency_key=%q",
			r.Method,
			r.URL.Path,
			time.Since(start).Round(time.Millisecond),
			r.Header.Get("Idempotency-Key"),
		)
	})
}
