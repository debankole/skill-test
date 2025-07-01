package service

import (
	"context"
	"report-service/client"
	"report-service/pdf"
)

type StudentReportService struct {
	studentClient client.StudentClient
}

func NewStudentReportService(studentClient client.StudentClient) *StudentReportService {
	return &StudentReportService{
		studentClient: studentClient,
	}
}

func (s *StudentReportService) GenerateStudentReportByID(ctx context.Context, id string) ([]byte, error) {
	student, err := s.studentClient.FetchStudent(ctx, id)
	if err != nil {
		return nil, err
	}
	return pdf.GenerateStudentReport(student)
}
