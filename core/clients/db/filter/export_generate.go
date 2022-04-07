package filter

import (
	"github.com/google/uuid"
	"github.com/kyaxcorp/go-core/core/helpers/file"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem"
	"github.com/xuri/excelize/v2"
)

func (e *Export) GeneratePdf() {

}

func (e *Export) GenerateExcel() bool {
	// based on the input Filter, we should query and generate

	// Let's initially export everything...
	f := excelize.NewFile()

	sheetName := "Sheet1"
	sheetIndex := f.GetSheetIndex(sheetName)

	for _, row := range e.items {
		for columnOrder, columnName := range e.columns {
			columnNr := columnOrder + 1
			if fieldValue, ok := row[columnName]; ok {
				axis, _err := excelize.ColumnNumberToName(columnNr)
				if _err != nil {
					e.excelError = _err
					return false
				}
				_err = f.SetCellValue(sheetName, axis, fieldValue)
				if _err != nil {
					e.excelError = _err
					return false
				}
			}
		}
	}

	// Generate a unique id
	id, _err := uuid.NewRandom()
	if _err != nil {
		e.excelError = _err
		return false
	}

	fileExtension := "xlsx"
	fileName := id.String()
	fullFileName := fileName + "." + fileExtension

	e.excelFileName = fileName
	e.excelFullFileName = fullFileName
	e.excelFileExtension = fileExtension

	tmpPath, _err := filesystem.GetAppTmpPath()
	if _err != nil {
		e.excelError = _err
		return false
	}
	fullFilePath := tmpPath + fileName

	// Set active sheet of the workbook.
	f.SetActiveSheet(sheetIndex)
	// Save spreadsheet by the given path.
	if err := f.SaveAs(fullFilePath); err != nil {
		e.excelError = _err
		return false
	}

	e.excelFileSizeBytes, _err = file.Size(fullFilePath)
	if _err != nil {
		e.excelError = _err
		return false
	}

	return true
}

func (e *Export) QueryItems() error {
	// based on the input Filter, we should query and generate

	dbResult := e.Filter.DB().Find(&e.items)

	if dbResult.Error != nil {
		return dbResult.Error
	}
	e.nrOfRows = int64(len(e.items))
	e.itemsSet = true
	return nil
}

func (e *Export) SetColumns(columns []string) *Export {
	e.columns = columns
	return e
}

func (e *Export) SetItems(items []map[string]interface{}) *Export {
	e.itemsSet = true
	e.items = items
	return e
}
