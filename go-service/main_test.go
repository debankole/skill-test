package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/ledongthuc/pdf"
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

	// --- PDF content check ---
	text, err := extractTextFromBytes(body)
	require.NoError(t, err)

	require.Contains(t, text, "John Doe")
	require.Contains(t, text, "john.doe@example.com")
	require.Contains(t, text, "10")
	require.Contains(t, text, "A")
}

func extractTextFromBytes(pdfData []byte) (string, error) {
	reader := bytes.NewReader(pdfData)
	pdfReader, err := pdf.NewReader(reader, int64(len(pdfData)))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	totalPage := pdfReader.NumPage()
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		page := pdfReader.Page(pageIndex)
		content, err := page.GetPlainText(nil)
		if err != nil {
			return "", err
		}
		buf.WriteString(content)
	}

	return buf.String(), nil
}
