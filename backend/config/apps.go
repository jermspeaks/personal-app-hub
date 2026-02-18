package config

import (
	"path/filepath"
)

// App represents metadata for a personal app
type App struct {
	Name        string
	Description string
	Path        string
	Ports       []int
}

// GetApps returns the configuration for all personal apps
func GetApps() []App {
	// Base path is one level up from hub directory
	basePath := filepath.Join("..", "..")
	
	return []App{
		{
			Name:        "audiophile",
			Description: "Audio/music related application",
			Path:        filepath.Join(basePath, "audiophile"),
			Ports:       []int{8000, 5173}, // backend, frontend
		},
		{
			Name:        "blippy",
			Description: "Blip cultivation: capture small sparks of attention, resurface them in a feed",
			Path:        filepath.Join(basePath, "blippy"),
			Ports:       []int{6900}, // Next.js single port
		},
		{
			Name:        "contacts-app",
			Description: "Contact management with import/export capabilities (Google Contacts, LinkedIn, Obsidian)",
			Path:        filepath.Join(basePath, "contacts-app"),
			Ports:       []int{4001, 4000}, // backend, frontend
		},
		{
			Name:        "digital-leatherman",
			Description: "Multi-purpose utility application",
			Path:        filepath.Join(basePath, "digital-leatherman"),
			Ports:       []int{8100, 5273}, // backend, frontend
		},
		{
			Name:        "slowtube",
			Description: "Slow media consumption app (YouTube, movies, TV shows)",
			Path:        filepath.Join(basePath, "slowtube"),
			Ports:       []int{3001, 5200}, // backend, frontend
		},
	}
}
