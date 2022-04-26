package util

import (
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx/v3"
	"path/filepath"
	"strings"
)

// 第一行注释，第二行程序用field，第三行类型 第四行数据

const tint = "int"
const tLong = "long"
const tFloat = "float"
const tJson = "json"
const tString = "string"
const tList = "list"

type ExcelData struct {
	Name   string
	Sheets []ExcelSheetData
}

type ExcelSheetData struct {
	Name string
	Rows []ExcelRowData
}

type ExcelRowData struct {
	DataMap map[string]interface{}
}

type FieldType struct {
	Field string
	Type  string
}

func ReadXlsx(filename string) []ExcelRowData {
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
		return nil
	}
	baseName := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
	excelData := ExcelData{Name: baseName}

	for _, sheet := range xlFile.Sheets {
		excelData.Sheets = append(excelData.Sheets, makeSheet(sheet))
	}

	return excelData.Sheets[0].Rows
}

func makeSheet(sheet *xlsx.Sheet) ExcelSheetData {
	s := ExcelSheetData{}
	s.Name = sheet.Name

	field, _ := sheet.Row(1)

	fieldType, _ := sheet.Row(2)

	m := make(map[int]FieldType)

	maxCol := sheet.MaxCol

	for i := 0; i < maxCol; i++ {
		ft := fieldType.GetCell(i)
		c := field.GetCell(i)
		v := FieldType{
			Field: strings.ReplaceAll(c.Value, "_", ""),
			Type:  ft.Value,
		}
		m[i] = v
	}

	// rows
	for i := 3; i < sheet.MaxRow; i++ {
		r, _ := sheet.Row(i)
		row := makeRow(i, r, m)
		if len(row.DataMap) == 0 {
			continue
		}
		s.Rows = append(s.Rows, row)
	}

	return s
}

func makeRow(rowNum int, row *xlsx.Row, m map[int]FieldType) ExcelRowData {

	rd := ExcelRowData{
		DataMap: make(map[string]interface{}),
	}
	colCount := len(m)

	if row.Sheet.MaxCol != len(m) {
		colCount = row.Sheet.MaxCol
	}

	for i := 0; i < colCount; i++ {
		var cell interface{}
		var e error

		if row.GetCell(i).Value == "" {
			continue
		}
		switch m[i].Type {
		case tint:
			cell, e = row.GetCell(i).Int()
		case tLong:
			cell, e = row.GetCell(i).Int64()
		case tFloat:
			cell, e = row.GetCell(i).Float()
		case tJson:
			e = json.Unmarshal([]byte(row.GetCell(i).Value), &cell)
		case tString, tList:
			cell = row.GetCell(i).String()
		default:
			cell = row.GetCell(i).Value
		}

		if e != nil {
			println(row.Sheet.Name, "行", rowNum, "列", i)
			panic(e)
		}
		rd.DataMap[m[i].Field] = cell
	}
	//delete(rd.DataMap, "")

	return rd
}
