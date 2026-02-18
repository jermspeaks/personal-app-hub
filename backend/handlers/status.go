package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"hub/backend/config"
	"hub/backend/services"
)

// AppStatus represents the status of an app
type AppStatus struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Online      bool      `json:"online"`
	LastUpdated time.Time `json:"lastUpdated"`
	Ports       []int     `json:"ports"`
}

// Status handles GET /api/status
func Status(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	apps := config.GetApps()
	statuses := make([]AppStatus, 0, len(apps))

	for _, app := range apps {
		// Check if app is online (at least one port is listening)
		online := services.CheckPorts(app.Ports)

		// Get last commit date
		lastUpdated := services.GetLastCommitDate(app.Path)

		statuses = append(statuses, AppStatus{
			Name:        app.Name,
			Description: app.Description,
			Online:      online,
			LastUpdated: lastUpdated,
			Ports:       app.Ports,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(statuses); err != nil {
		slog.Error("failed to encode response", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
