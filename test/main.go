package main

import (
	"fmt"

	"github.com/esourceful/xlsxreader"
)

func main() {
	filepath := "./sample_parse_failure.xlsx"
	//filepath := "./test-small.xlsx"
	e, err := xlsxreader.OpenFile(filepath)
	if err != nil {
		fmt.Printf("error: %s \n", err)
		return
	}
	defer e.Close()

	fmt.Printf("Worksheets: %s \n", e.Sheets)

	for row := range e.ReadRows(e.Sheets[0]) {
		if row.Error != nil {
			fmt.Printf("error on row %d: %s \n", row.Index, row.Error)
			return
		}

		if row.Index < 10 {
			fmt.Printf("%+v \n", row.Cells)
		}
	}
}
