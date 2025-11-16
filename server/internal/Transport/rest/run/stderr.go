package run

import "strings"

// helper to attach stderr if present
func maybeStderr(stderr string) string {
	stderr = strings.TrimSpace(stderr)
	if stderr == "" {
		return ""
	}
	return "\n[stderr] " + stderr
}
