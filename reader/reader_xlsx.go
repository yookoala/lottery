package reader

import (
	"github.com/tealeg/xlsx"

	"fmt"
)

func OpenXLSXSheet(filename string, sheetNum int) (c Collection, err error) {
	f, err := xlsx.OpenFile(filename)
	if err != nil {
		return
	}

	if len(f.Sheets) <= sheetNum {
		err = fmt.Errorf("sheet %d not exists in file %#v", sheetNum, filename)
		return
	}

	c = &XLSXCollection{f.Sheets[sheetNum]}
	return
}

type XLSXCollection struct {
	sheet *xlsx.Sheet
}

func (c *XLSXCollection) Len() int {
	return len(c.sheet.Rows)
}

func (c *XLSXCollection) Row(rowNum int) (r Row, err error) {
	if c.Len() <= rowNum {
		err = fmt.Errorf("row %d not exists in sheet", rowNum)
		return
	}
	r = &XLSXRow{c.sheet.Rows[rowNum]}
	return
}

type XLSXRow struct {
	row *xlsx.Row
}

func (r *XLSXRow) Len() int {
	return len(r.row.Cells)
}

func (r *XLSXRow) ReadString(cellNum int) string {
	if r.Len() <= cellNum {
		return ""
	}
	return r.row.Cells[cellNum].Value
}
