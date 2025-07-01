package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"

	"report-service/client"
	"report-service/service"
)

type StudentReportHandler struct {
	client client.StudentClient
}

func NewStudentReportHandler(client client.StudentClient) *StudentReportHandler {
	return &StudentReportHandler{client: client}
}

func (h *StudentReportHandler) StudentReportHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing student ID", http.StatusBadRequest)
		return
	}

	pdfBytes, err := service.NewStudentReportService(h.client).GenerateStudentReportByID(r.Context(), id)
	if err != nil {
		log.Errorf("Failed to generate student report: %v", err)
		http.Error(w, "Failed to generate student report", http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=student_report.pdf")
	w.Write(pdfBytes)
}
