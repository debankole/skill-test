package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Student struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	Email              string  `json:"email"`
	SystemAccess       bool    `json:"systemAccess"`
	Phone              *string `json:"phone"`
	Gender             *string `json:"gender"`
	Dob                *string `json:"dob"`
	Class              *string `json:"class"`
	Section            *string `json:"section"`
	Roll               *string `json:"roll"`
	FatherName         *string `json:"fatherName"`
	FatherPhone        *string `json:"fatherPhone"`
	MotherName         *string `json:"motherName"`
	MotherPhone        *string `json:"motherPhone"`
	GuardianName       *string `json:"guardianName"`
	GuardianPhone      *string `json:"guardianPhone"`
	RelationOfGuardian *string `json:"relationOfGuardian"`
	CurrentAddress     *string `json:"currentAddress"`
	PermanentAddress   *string `json:"permanentAddress"`
	AdmissionDate      *string `json:"admissionDate"`
	ReporterName       *string `json:"reporterName"`
}

type StudentClient interface {
	FetchStudent(ctx context.Context, id string) (*Student, error)
}

type NodeBackendClient struct {
	backendURL string
}

func NewNodeBackendClient() *NodeBackendClient {
	backendURL := os.Getenv("NODE_BACKEND_URL")
	if backendURL == "" {
		panic("NODE_BACKEND_URL is not set")
	}
	return &NodeBackendClient{
		backendURL: backendURL,
	}
}

func (c *NodeBackendClient) FetchStudent(ctx context.Context, id string) (*Student, error) {
	url := fmt.Sprintf("%s/api/v1/students/%s", c.backendURL, id)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("backend returned status %d", resp.StatusCode)
	}

	var student Student
	if err := json.NewDecoder(resp.Body).Decode(&student); err != nil {
		return nil, err
	}
	return &student, nil
}
