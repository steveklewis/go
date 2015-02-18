package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

// Write a file, then read from in-memory
func main() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet = file.AddSheet("Sheet1")
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "I am a cell!"

	for _, tempSheet := range file.Sheets {
		for _, tempRow := range tempSheet.Rows {
			for _, tempCell := range tempRow.Cells {
				fmt.Printf("%s\n", tempCell.String())
			}
		}
	}

	err = file.Save("FirstExcel.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}

	file.Sheets[0].Rows[0].Cells[0].Value = "I am a cell2!"

	for _, tempSheet := range file.Sheets {
		for _, tempRow := range tempSheet.Rows {
			for _, tempCell := range tempRow.Cells {
				fmt.Printf("%s\n", tempCell.String())
			}
		}
	}

	err = file.Save("SecondExcel.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
