package pdf

import (
	"log"

	fillpdf "github.com/goevexx/fillpdfV2"
)

func FillPDF(form fillpdf.Form, tmplPath, outPath string, useAnsi bool) {
	// Fill the form PDF with our values.
	err := fillpdf.Fill(form, tmplPath, outPath, useAnsi, true)
	if err != nil {
		log.Fatal(err)
	}
}
