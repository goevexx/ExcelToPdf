package excel

import (
	"log"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

func ParseExcel(excelFileName, dateFormat string) (result [][]string) {
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}
	result = [][]string{}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			rowArr := []string{}
			for _, cell := range row.Cells {
				if cell.String() != "" && strings.ContainsAny(cell.GetNumberFormat(), "dmy") && cell.Type() == xlsx.CellTypeNumeric { //TODO: continue
					cFormat := strings.ToLower(cell.GetNumberFormat())
					cFormat = strings.Replace(cFormat, "dd", "02", 1)
					cFormat = strings.Replace(cFormat, "d", "2", 1)
					cFormat = strings.Replace(cFormat, "mm", "01", 1)
					cFormat = strings.Replace(cFormat, "m", "1", 1)
					cFormat = strings.Replace(cFormat, "yyyy", "2006", 1)
					cFormat = strings.Replace(cFormat, "yy", "06", 1)
					t, err := time.Parse(cFormat, cell.String())
					if err != nil {
						log.Println(err)
						continue
					}
					rowArr = append(rowArr, t.Format(dateFormat))
				} else {
					rowArr = append(rowArr, cell.String())
				}
			}
			result = append(result, rowArr)
		}
	}
	return
}
