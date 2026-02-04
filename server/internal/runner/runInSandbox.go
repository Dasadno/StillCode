package runner

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// Resource limits for Docker containers
const (
	MemoryLimit  = "128m"
	CPULimit     = "0.5"
	PidsLimit    = "50"
	TimeoutSec   = 5
	MaxCodeSize  = 1024 * 1024 // 1MB
	MaxInputSize = 1024 * 100  // 100KB
)

func RunInSandbox(ctx context.Context, language, code, input string) (stdout, stderr string, durationMs int64, status string, err error) {
	// Validate input sizes
	if len(code) > MaxCodeSize {
		return "", "", 0, "error", fmt.Errorf("code size exceeds limit (%d bytes)", MaxCodeSize)
	}
	if len(input) > MaxInputSize {
		return "", "", 0, "error", fmt.Errorf("input size exceeds limit (%d bytes)", MaxInputSize)
	}

	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "sc_runner_*")
	if err != nil {
		return "", "", 0, "error", fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer os.RemoveAll(tempDir)

	var (
		filename   string
		image      string
		compileCmd []string
		runCmd     []string
	)

	// Language configuration
	switch strings.ToLower(language) {
	case "python", "py":
		filename = "main.py"
		image = "python:3.11-slim"
		runCmd = []string{"python3", filename}

	case "cpp", "c++":
		filename = "main.cpp"
		image = "gcc:12"
		compileCmd = []string{"g++", "-O2", "-std=c++17", "-o", "main", "main.cpp"}
		runCmd = []string{"./main"}

	case "java":
		filename = "Main.java"
		image = "openjdk:20-slim"
		compileCmd = []string{"javac", "Main.java"}
		runCmd = []string{"java", "Main"}

	case "javascript", "js":
		filename = "main.js"
		image = "node:20-slim"
		runCmd = []string{"node", filename}

	case "go", "golang":
		filename = "main.go"
		image = "golang:1.23-alpine"
		runCmd = []string{"go", "run", filename}

	default:
		return "", "", 0, "error", fmt.Errorf("unsupported language: %s", language)
	}

	// Write source code
	sourcePath := filepath.Join(tempDir, filename)
	if err := os.WriteFile(sourcePath, []byte(code), 0644); err != nil {
		return "", "", 0, "error", fmt.Errorf("failed to write code: %w", err)
	}

	// Write input file if provided
	if input != "" {
		inputPath := filepath.Join(tempDir, "input.txt")
		if err := os.WriteFile(inputPath, []byte(input), 0644); err != nil {
			return "", "", 0, "error", fmt.Errorf("failed to write input: %w", err)
		}
	}

	// Docker arguments with resource limits
	dockerArgs := []string{
		"run",
		"--rm",
		"-i", // Enable stdin
		"--network", "none",
		"--memory", MemoryLimit,
		"--memory-swap", MemoryLimit,
		"--cpus", CPULimit,
		"--pids-limit", PidsLimit,
		"--read-only",
		"--tmpfs", "/tmp:size=64m,mode=1777",
		"-v", tempDir + ":/app:rw",
		"-w", "/app",
		image,
	}

	// Compilation step (if needed)
	if len(compileCmd) > 0 {
		compileArgs := append(dockerArgs, compileCmd...)
		compileCtx, compileCancel := context.WithTimeout(ctx, 30*time.Second)
		defer compileCancel()

		cmd := exec.CommandContext(compileCtx, "docker", compileArgs...)
		var compileOut, compileErr bytes.Buffer
		cmd.Stdout = &compileOut
		cmd.Stderr = &compileErr

		if err := cmd.Run(); err != nil {
			stderr = compileErr.String()
			if compileCtx.Err() == context.DeadlineExceeded {
				return "", "Compilation timeout", 0, "compile_error", nil
			}
			return "", stderr, 0, "compile_error", nil
		}
	}

	// Run the program
	runArgs := append(dockerArgs, runCmd...)

	runCtx, runCancel := context.WithTimeout(ctx, TimeoutSec*time.Second)
	defer runCancel()

	cmd := exec.CommandContext(runCtx, "docker", runArgs...)

	// Pipe input if provided
	if input != "" {
		cmd.Stdin = strings.NewReader(input)
	}

	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	start := time.Now()
	runErr := cmd.Run()
	durationMs = time.Since(start).Milliseconds()

	stdout = outBuf.String()
	stderr = errBuf.String()

	// Determine execution status
	if runCtx.Err() == context.DeadlineExceeded {
		return stdout, stderr, durationMs, "timeout", nil
	}
	if runErr != nil {
		// Check for memory limit exceeded
		if strings.Contains(stderr, "out of memory") || strings.Contains(stderr, "OOM") {
			return stdout, stderr, durationMs, "memory_limit", nil
		}
		return stdout, stderr, durationMs, "runtime_error", nil
	}

	return stdout, stderr, durationMs, "ok", nil
}
