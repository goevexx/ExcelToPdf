package main

import (
	"flag"
	"log"
	"path"
	"path/filepath"
	"strconv"
	"tools/exceltopdf/excel"
	"tools/exceltopdf/pdf"

	fillpdf "github.com/goevexx/fillpdfV2"
)

var (
	tmplPath   string
	outPath    string
	tablePath  string
	dateFormat string
	useAnsi    bool
)

func init() {
	flag.StringVar(&tmplPath, "pdf", "tpl.pdf", "Path to PDF file template")
	flag.StringVar(&outPath, "out", "./", "PDF file output path")
	flag.StringVar(&tablePath, "xlsx", "excel.xlsx", "EXCEL file path")
	flag.StringVar(&dateFormat, "dateformat", "02.01.2006", "Dateformat in PDF (with go's magic dateformat)")
	flag.BoolVar(&useAnsi, "ansi", false, "Use ANSI(Windows1250) charset for FDF file (solves umlauts in form field names)")

}

func main() {
	flag.Parse()

	log.Println("Parse xlsx")
	table := excel.ParseExcel(tablePath, dateFormat)
	if len(table) <= 0 {
		log.Fatal("Empty table.")
	}

	tableHead := table[0]

	for i := 1; i < len(table); i++ {
		log.Println("Check row " + strconv.Itoa(i))
		form := FormHolder{}
		form.Fill(tableHead, table[i])
		if !form.IsEmpty() {
			filePath := path.Dir(outPath) + string(filepath.Separator) + strconv.Itoa(i) + "_" + path.Base(outPath)
			log.Println("Create " + filePath)
			pdf.FillPDF(fillpdf.Form(form.Map), tmplPath, filePath, useAnsi)
		}
	}
}
