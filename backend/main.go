package main

import (
	"log/slog"
	"net/http"
	"os"
	"strings"

	"hub/backend/handlers"
	"hub/backend/middleware"
)

// initLogging configures structured logging based on environment variables.
// LOG_FORMAT: "json" (default) or "text" for human-readable output
// LOG_LEVEL: "debug", "info" (default), "warn", or "error"
func initLogging() {
	format := os.Getenv("LOG_FORMAT")
	if format == "" {
		format = "json" // Default to JSON for production
	}

	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level: parseLogLevel(os.Getenv("LOG_LEVEL")),
	}

	switch format {
	case "text":
		handler = slog.NewTextHandler(os.Stderr, opts)
	default:
		handler = slog.NewJSONHandler(os.Stderr, opts)
	}

	slog.SetDefault(slog.New(handler))
}

// parseLogLevel converts a string log level to slog.Level.
// Returns slog.LevelInfo as default if the level is invalid.
func parseLogLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func main() {
	// Initialize structured logging before any log calls
	initLogging()

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
		slog.Error("server failed", "err", err, "addr", addr)
		os.Exit(1)
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
