package plt4m

import "github.com/tealeg/xlsx"

// SchoolList Xlsx document representation
type SchoolList struct {
	*xlsx.File
}

// SchoolItem represents one row
// of the xlsx document
type SchoolItem struct {
	*xlsx.Row
}

// IterateCallback type to use
// in IterateOver
type IterateCallback func(int, *SchoolItem)
