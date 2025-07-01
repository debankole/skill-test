package client

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed testdata/student_example.json
var studentExampleJSON []byte

type MockStudentClient struct {
	StudentData *Student
	Err         error
}

func (m *MockStudentClient) FetchStudent(ctx context.Context, id string) (*Student, error) {
	return m.StudentData, m.Err
}

// NewMockStudentClientWithExampleData returns a mock client with example student data loaded from embedded JSON
func NewMockStudentClientWithExampleData() *MockStudentClient {
	var student Student
	err := json.Unmarshal(studentExampleJSON, &student)
	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal embedded student example JSON: %v", err))
	}
	return &MockStudentClient{
		StudentData: &student,
		Err:         nil,
	}
}
