package run

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

func runInSandbox(ctx context.Context, language, code, input string) (stdout, stderr string, durationMs int64, status string, err error) {
	// choose file name and docker image/compile/run commands
	var (
		filename   string
		image      string
		compileCmd []string
		runCmd     []string
	)

	switch language {
	case "python":
		filename = "main.py"
		image = "python:3.11-slim"
		// Run python directly
		runCmd = []string{"python", filename}
	case "cpp", "c++":
		filename = "main.cpp"
		image = "gcc:12"
		compileCmd = []string{"bash", "-lc", "g++ -O2 -std=c++17 -o main main.cpp 2>compile_err.txt || true"}
		runCmd = []string{"bash", "-lc", "./main"}
	case "java":
		filename = "Main.java"
		image = "openjdk:20"
		compileCmd = []string{"bash", "-lc", "javac Main.java 2>compile_err.txt || true"}
		runCmd = []string{"bash", "-lc", "java Main"}
	default:
		return "", "unsupported language", 0, "error", fmt.Errorf("unsupported language")
	}

	script := &bytes.Buffer{}
	// write file
	fmt.Fprintf(script, "cat > %s <<'EOF'\n%s\nEOF\n", filename, code)
	// compile if needed
	if len(compileCmd) > 0 {
		fmt.Fprintln(script, compileCmd[0])
		for _, part := range compileCmd[1:] {
			fmt.Fprintln(script, part)
		}
	}

	fmt.Fprintln(script, "timeout 5s "+runCmd[0])

	dockerArgs := []string{"run", "--rm", "--network", "none", "-i", image, "bash", "-lc", script.String()}

	// Prepare command with context (timeout for whole container)
	overallTimeout := 6 * time.Second
	runCtx, cancel := context.WithTimeout(ctx, overallTimeout)
	defer cancel()

	cmd := exec.CommandContext(runCtx, "docker", dockerArgs...)
	// pass input to process (none needed)
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	start := time.Now()
	err = cmd.Run()
	duration := time.Since(start)
	durationMs = duration.Milliseconds()

	stdout = outBuf.String()
	stderr = errBuf.String()

	if runCtx.Err() == context.DeadlineExceeded {
		status = "timeout"
		return
	}
	if err != nil {
		// if docker failed, surface error
		status = "runtime_error"
		// but still return stdout/stderr
		return
	}
	status = "ok"
	return
}
