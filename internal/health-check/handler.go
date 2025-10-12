// Package healthcheck provides a simple health check handler for HTTP servers.
package healthcheck

import (
	"net/http"

	"ecommerce-service/pkg/httpx"
)

type HealthCheckHandler struct{}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (hc *HealthCheckHandler) Check(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	httpx.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
