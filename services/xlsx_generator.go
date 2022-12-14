package services

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func Generator() {

	f := excelize.NewFile()
	// Create a new worksheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set the active worksheet of the workbook.
	f.SetActiveSheet(index)
	// Save the spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

}
