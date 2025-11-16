package runner

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func runInSandbox(ctx context.Context, language, code, input string) (stdout, stderr string, durationMs int64, status string, err error) {
	// 1) Создаем временную директорию
	tempDir, err := os.MkdirTemp("", "sc_runner_*")
	if err != nil {
		return "", "", 0, "error", fmt.Errorf("failed to create temp dir: %w", err)
	}
	// Удаляем после выполнения
	defer os.RemoveAll(tempDir)

	var (
		filename   string
		image      string
		compileCmd []string
		runCmd     []string
	)

	// 2) Определяем настройки по языку
	switch language {
	case "python":
		filename = "main.py"
		image = "python:3.11-slim"
		runCmd = []string{"python3", filename}

	case "cpp", "c++":
		filename = "main.cpp"
		image = "gcc:12"
		compileCmd = []string{"g++", "-O2", "-std=c++17", "main.cpp", "-o", "main"}
		runCmd = []string{"./main"}

	case "java":
		filename = "Main.java"
		image = "openjdk:20"
		compileCmd = []string{"javac", "Main.java"}
		runCmd = []string{"java", "Main"}

	default:
		return "", "", 0, "error", fmt.Errorf("unsupported language: %s", language)
	}

	// 3) Создаем файлы в tempDir
	sourcePath := filepath.Join(tempDir, filename)
	if err := os.WriteFile(sourcePath, []byte(code), 0644); err != nil {
		return "", "", 0, "error", fmt.Errorf("failed to write code: %w", err)
	}

	if input != "" {
		if err := os.WriteFile(filepath.Join(tempDir, "input.txt"), []byte(input), 0644); err != nil {
			return "", "", 0, "error", fmt.Errorf("failed to write input: %w", err)
		}
	}

	// 4) Docker args
	dockerArgs := []string{
		"run",
		"--rm",
		"--network", "none",
		"-v", tempDir + ":/app",
		"-w", "/app",
		image,
	}

	// 5) Компиляция (если нужна)
	if len(compileCmd) > 0 {
		compileArgs := append(dockerArgs, compileCmd...)
		cmd := exec.CommandContext(ctx, "docker", compileArgs...)

		var compileErr bytes.Buffer
		cmd.Stderr = &compileErr

		if err := cmd.Run(); err != nil {
			stderr = compileErr.String()
			return "", stderr, 0, "compile_error", nil
		}
	}

	// 6) Запуск программы
	runArgs := append(dockerArgs, runCmd...)

	runCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(runCtx, "docker", runArgs...)

	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	start := time.Now()
	err = cmd.Run()
	durationMs = time.Since(start).Milliseconds()

	stdout = outBuf.String()
	stderr = errBuf.String()

	// 7) Обработка статуса выполнения
	if runCtx.Err() == context.DeadlineExceeded {
		return stdout, stderr, durationMs, "timeout", nil
	}
	if err != nil {
		return stdout, stderr, durationMs, "runtime_error", nil
	}

	return stdout, stderr, durationMs, "ok", nil
}
