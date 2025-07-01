package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"

	"report-service/service"
)

func StudentReportHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing student ID", http.StatusBadRequest)
		return
	}

	pdfBytes, err := service.GenerateStudentReportByID(r.Context(), id)
	if err != nil {
		log.Errorf("Failed to generate student report: %v", err)
		http.Error(w, "Failed to generate student report", http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=student_report.pdf")
	w.Write(pdfBytes)
}
