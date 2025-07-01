package service

import (
	"context"
	"report-service/client"
	"report-service/pdf"
)

func GenerateStudentReportByID(ctx context.Context, id string) ([]byte, error) {
	student, err := client.FetchStudent(ctx, id)
	if err != nil {
		return nil, err
	}
	return pdf.GenerateStudentReport(student)
}
