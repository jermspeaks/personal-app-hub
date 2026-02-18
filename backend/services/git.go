package services

import (
	"log/slog"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// GetLastCommitDate gets the last commit date from a git repository
// Returns the commit time, or zero time if there's an error
func GetLastCommitDate(repoPath string) time.Time {
	// Resolve absolute path
	absPath, err := filepath.Abs(repoPath)
	if err != nil {
		slog.Warn("failed to resolve path", "path", repoPath, "err", err)
		return time.Time{}
	}

	// Run git command to get last commit timestamp
	cmd := exec.Command("git", "-C", absPath, "log", "-1", "--format=%ct", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		slog.Debug("git command failed", "path", absPath, "err", err)
		return time.Time{}
	}

	// Parse Unix timestamp
	timestampStr := string(output)
	if len(timestampStr) == 0 {
		return time.Time{}
	}

	// Trim whitespace (including newline)
	timestampStr = strings.TrimSpace(timestampStr)
	if len(timestampStr) == 0 {
		return time.Time{}
	}

	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		slog.Warn("failed to parse timestamp", "timestamp", timestampStr, "err", err)
		return time.Time{}
	}

	return time.Unix(timestamp, 0)
}
