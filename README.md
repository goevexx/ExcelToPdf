# ExcelToPdf

Handy tool, to insert values from an excel sheet into a predefined pdf template fields.

## Prequisites

[PDFToolkit](https://www.pdflabs.com/tools/pdftk-server/) needs to be installed.

## Usage
```
Usage of exceltopdf.exe:
  -ansi
        Use ANSI(Windows1250) charset for FDF file (solves umlauts in form field names)
  -dateformat string
        Dateformat in PDF (with go's magic dateformat) (default "02.01.2006")
  -out string
        PDF file output path (default "./")
  -pdf string
        Path to PDF file template (default "tpl.pdf")
  -xlsx string
        EXCEL file path (default "excel.xlsx")
```

### Example
```
exceltopdf -pdf "test/Ausbildungsnachweis_tmpl.pdf" \
-xlsx "test/ausbildungsnachweise1.xlsx" \
-out "test/out/Ausbildungsnachweis.pdf" \
-dateformat 02.01 \
-ansi 
```

## FAQ
### Cell from my XLSX sheet is not being pasted into the PDF?
Make sure that the field names in your XLSX sheet fit the field names in your pdf.
In case this doesn't work, try to use the `-ansi` flag.  

### Further problems?

Create an issue.
