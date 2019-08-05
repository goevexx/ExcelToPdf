package excel_test

import (
	"testing"
	. "tools/exceltopdf/excel"
)

func TestParseExcel(t *testing.T) {
	t.Log(ParseExcel("tables/test.xlsx"))
}
