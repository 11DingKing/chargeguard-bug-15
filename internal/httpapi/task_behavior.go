package httpapi

import (
	"chargeguard/internal/charging"
	"encoding/json"
	"net/http"
)

var hazardRepository = charging.NewHazardRepository()

func TaskHTTPHandler(w http.ResponseWriter, r *http.Request) {
	tenant := r.URL.Query().Get("tenant")
	status := r.URL.Query().Get("status")
	if tenant == "" || status == "" {
		http.Error(w, "tenant and status required", http.StatusBadRequest)
		return
	}
	items := hazardRepository.List(tenant, status)
	total := hazardRepository.Count(tenant, status)
	if len(items) != total {
		http.Error(w, "inconsistent tenant result", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"items": items, "total": total})
}
