package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"hub/backend/handlers"
	"hub/backend/middleware"
)

func main() {
	mux := http.NewServeMux()

	// Register status endpoint
	mux.HandleFunc("/api/status", cors(handlers.Status))

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8518"
	}
	addr := ":" + port

	// Apply middleware: recovery -> logging -> mux
	handler := middleware.Recovery(middleware.Logging(mux))

	slog.Info("server listening", "addr", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}

// cors wraps a handler to add CORS headers for the frontend dev server.
// In development, the request Origin is reflected so any localhost port works.
func cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		// Allow any localhost origin in development (e.g. 8517, 5173, 3000).
		if origin != "" && (strings.HasPrefix(origin, "http://localhost:") || strings.HasPrefix(origin, "http://127.0.0.1:")) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}
