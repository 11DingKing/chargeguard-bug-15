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
	_ = json.NewEncoder(w).Encode(map[string]any{"items": hazardRepository.List(tenant, status), "total": hazardRepository.Count(tenant, status)})
}
