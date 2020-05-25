package plt4m

import (
	"fmt"
	"os"
	"path"

	"github.com/tealeg/xlsx"
)

// GetStateID returns state ID of the row
func (item *SchoolItem) GetStateID() int {
	id, _ := item.Cells[2].Int()
	return id
}

// GetSchoolName returns state ID of the row
func (item *SchoolItem) GetSchoolName() string {
	return item.Cells[1].String()
}

// AddOffset writes value to the offset cell
func (item *SchoolItem) AddOffset(offset string) {
	cell := item.AddCell()
	cell.Value = offset
}

// Mark colors background of the row RED
func (item *SchoolItem) Mark() {
	for i := 0; i < 4; i++ {
		style := xlsx.NewStyle()
		style.Fill = *xlsx.FillRed
		style.Alignment = xlsx.Alignment{
			Horizontal: "left",
		}
		item.Cells[i].SetStyle(style)
	}
}

// Sheet returns the main working document sheet
func (file *SchoolList) Sheet() *xlsx.Sheet {
	sheet := file.Sheets[0]
	if sheet == nil {
		panic("Sheet not found")
	}
	return sheet
}

// PrintCount writes stdout Rows count
// Of the document
func (file *SchoolList) PrintCount() {
	sheet := file.Sheet()
	fmt.Printf("Schools count %d\n", len(sheet.Rows))
}

// AddColumn adds new column on 3rd cell
// of the first row
func (file *SchoolList) AddColumn(name string) {
	sheet := file.Sheet()
	newCell := sheet.Rows[0].AddCell()
	style := &xlsx.Style{}
	style.Font.Underline = true
	style.Font.Bold = true

	newCell.SetStyle(style)
	newCell.Value = name
}

// IterateOver iterates over all rows
// of the document invoking callback with the
// current row and row index
func (file *SchoolList) IterateOver(callback IterateCallback) {
	sheet := file.Sheet()
	for index, row := range sheet.Rows {
		if index == 0 {
			continue
		}

		callback(index, &SchoolItem{row})
	}
}

// WriteUpdates saves all changes to the file
func (file *SchoolList) WriteUpdates(outputFilename string) error {
	dir, _ := os.Getwd()
	f, err := os.Create(path.Join(dir, "edited.xlsx"))
	defer f.Close()

	if err != nil {
		return err
	}

	err = file.Write(f)
	if err != nil {
		return err
	}
	return nil
}
