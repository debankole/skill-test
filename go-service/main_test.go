package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestIntegration_StudentReportPDF(t *testing.T) {
	// Start the Go service as a subprocess
	cmd := exec.Command("go", "run", "main.go")
	cmd.Env = append(os.Environ(), "PORT=18080", "USE_MOCK=true", "NODE_BACKEND_URL=http://localhost:3000")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	require.NoError(t, err)
	defer func() {
		_ = cmd.Process.Kill()
	}()

	// Wait for the server to be ready
	time.Sleep(2 * time.Second)

	resp, err := http.Get("http://localhost:18080/api/v1/students/1/report")
	require.NoError(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, "application/pdf", resp.Header.Get("Content-Type"))
	require.True(t, len(body) > 100, "PDF body should not be empty")
}
