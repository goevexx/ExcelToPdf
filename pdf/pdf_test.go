package pdf_test

import (
	"testing"
	. "tools/exceltopdf/pdf"

	"github.com/desertbit/fillpdf"
)

func TestFillPDF(t *testing.T) {
	form := fillpdf.Form{
		"Ausbildungsjahr":           "2019",
		"Ggf ausbildende Abteilung": "Test",
		"Unterweisungen, betrieblicher Unterricht, sonstige Schulungen": "Test\nTest1",
	}
	FillPDF(form, "files/Ausbildungsnachweis_tmpl.pdf", "files/filled.pdf")
}
