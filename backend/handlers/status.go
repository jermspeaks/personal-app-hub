package handlers

import (
	"encoding/json"
	"fmt"
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
	URL         string    `json:"url"`
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

		// Determine frontend port (use last port in array, typically the frontend)
		var frontendPort int
		if len(app.Ports) > 0 {
			frontendPort = app.Ports[len(app.Ports)-1]
		}
		url := fmt.Sprintf("http://localhost:%d", frontendPort)

		statuses = append(statuses, AppStatus{
			Name:        app.Name,
			Description: app.Description,
			Online:      online,
			LastUpdated: lastUpdated,
			Ports:       app.Ports,
			URL:         url,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(statuses); err != nil {
		slog.Error("failed to encode response", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
