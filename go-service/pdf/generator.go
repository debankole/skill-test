package pdf

import (
	"bytes"
	"report-service/client"

	"github.com/phpdave11/gofpdf"
)

func GenerateStudentReport(student *client.Student) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Student Report")
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Name: "+student.Name)
	pdf.Ln(8)
	pdf.Cell(40, 10, "Email: "+student.Email)
	pdf.Ln(8)
	if student.Class != nil {
		pdf.Cell(40, 10, "Class: "+*student.Class)
		pdf.Ln(8)
	}
	if student.Section != nil {
		pdf.Cell(40, 10, "Section: "+*student.Section)
		pdf.Ln(8)
	}
	if student.Roll != nil {
		pdf.Cell(40, 10, "Roll: "+*student.Roll)
		pdf.Ln(8)
	}
	if student.Phone != nil {
		pdf.Cell(40, 10, "Phone: "+*student.Phone)
		pdf.Ln(8)
	}
	if student.Gender != nil {
		pdf.Cell(40, 10, "Gender: "+*student.Gender)
		pdf.Ln(8)
	}
	if student.Dob != nil {
		pdf.Cell(40, 10, "DOB: "+*student.Dob)
		pdf.Ln(8)
	}
	if student.FatherName != nil {
		pdf.Cell(40, 10, "Father's Name: "+*student.FatherName)
		pdf.Ln(8)
	}
	if student.FatherPhone != nil {
		pdf.Cell(40, 10, "Father's Phone: "+*student.FatherPhone)
		pdf.Ln(8)
	}
	if student.MotherName != nil {
		pdf.Cell(40, 10, "Mother's Name: "+*student.MotherName)
		pdf.Ln(8)
	}
	if student.MotherPhone != nil {
		pdf.Cell(40, 10, "Mother's Phone: "+*student.MotherPhone)
		pdf.Ln(8)
	}
	if student.GuardianName != nil {
		pdf.Cell(40, 10, "Guardian's Name: "+*student.GuardianName)
		pdf.Ln(8)
	}
	if student.GuardianPhone != nil {
		pdf.Cell(40, 10, "Guardian's Phone: "+*student.GuardianPhone)
		pdf.Ln(8)
	}
	if student.RelationOfGuardian != nil {
		pdf.Cell(40, 10, "Relation of Guardian: "+*student.RelationOfGuardian)
		pdf.Ln(8)
	}
	if student.CurrentAddress != nil {
		pdf.Cell(40, 10, "Current Address: "+*student.CurrentAddress)
		pdf.Ln(8)
	}
	if student.PermanentAddress != nil {
		pdf.Cell(40, 10, "Permanent Address: "+*student.PermanentAddress)
		pdf.Ln(8)
	}
	if student.AdmissionDate != nil {
		pdf.Cell(40, 10, "Admission Date: "+*student.AdmissionDate)
		pdf.Ln(8)
	}
	if student.ReporterName != nil {
		pdf.Cell(40, 10, "Reporter: "+*student.ReporterName)
		pdf.Ln(8)
	}
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	return buf.Bytes(), err
}
