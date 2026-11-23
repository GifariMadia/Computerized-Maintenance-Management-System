package response

import (
	"encoding/json"
	"net/http"
	"time"
)

// Meta conveys WIB server timestamp, pagination, and data freshness (URS 2.2c).
type Meta struct {
	TimestampWIB string `json:"timestamp_wib"`
	DataStale    bool   `json:"data_stale,omitempty"`
	StaleMessage string `json:"stale_message,omitempty"`
}

type Envelope struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   *Error `json:"error,omitempty"`
	Meta    Meta   `json:"meta"`
}

type Error struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields,omitempty"`
}

func nowWIB() string {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		loc = time.FixedZone("WIB", 7*3600)
	}
	return time.Now().In(loc).Format(time.RFC3339)
}

func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(Envelope{
		Success: status >= 200 && status < 300,
		Data:    data,
		Meta: Meta{
			TimestampWIB: nowWIB(),
		},
	})
}

func Fail(w http.ResponseWriter, status int, code, message string, fields map[string]string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(Envelope{
		Success: false,
		Error: &Error{
			Code:    code,
			Message: message,
			Fields:  fields,
		},
		Meta: Meta{
			TimestampWIB: nowWIB(),
		},
	})
}
