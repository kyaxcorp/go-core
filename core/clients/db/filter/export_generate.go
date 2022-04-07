package filter

import (
	"github.com/google/uuid"
	"github.com/kyaxcorp/go-core/core/helpers/conv"
	"github.com/kyaxcorp/go-core/core/helpers/file"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem"
	"github.com/xuri/excelize/v2"
	"time"
)

func (e *Export) GeneratePdf() {

}

func (e *Export) GetExportPath(paths ...string) (string, error) {
	_paths := append([]string{"exporter"}, paths...)
	tmpPath, _err := filesystem.GetAppTmpPath(_paths...)
	if _err != nil {
		return "", _err
	}
	return tmpPath, nil
}

func (e *Export) GetExcelExportPath() (string, error) {
	tmpPath, _err := e.GetExportPath("excel")
	if _err != nil {
		return "", _err
	}
	return tmpPath, nil
}

func (e *Export) GetPdfExportPath() (string, error) {
	tmpPath, _err := e.GetExportPath("pdf")
	if _err != nil {
		return "", _err
	}
	return tmpPath, nil
}

func (e *Export) GetHtmlExportPath() (string, error) {
	tmpPath, _err := e.GetExportPath("html")
	if _err != nil {
		return "", _err
	}
	return tmpPath, nil
}

func (e *Export) GetJsonExportPath() (string, error) {
	tmpPath, _err := e.GetExportPath("json")
	if _err != nil {
		return "", _err
	}
	return tmpPath, nil
}

func (e *Export) GetWordExportPath() (string, error) {
	tmpPath, _err := e.GetExportPath("word")
	if _err != nil {
		return "", _err
	}
	return tmpPath, nil
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

	// TODO: we should save the file id into a memory stack with id's... maybe in the filter somewhere...

	if e.ExportName == "" {
		e.ExportName = id.String()
	}

	now := time.Now()
	e.excelFileID = id
	e.excelFileName = e.ExportName + "_" + conv.Int64ToStr(now.UnixMilli())
	e.excelFullFileName = fullFileName
	e.excelFileExtension = fileExtension

	tmpPath, _err := e.GetExcelExportPath()
	if _err != nil {
		e.excelError = _err
		return false
	}

	fullFilePath := tmpPath + fileName
	e.excelFullFilePath = fullFilePath

	// Set active sheet of the workbook.
	f.SetActiveSheet(sheetIndex)
	// Save spreadsheet by the given path.
	if err := f.SaveAs(fullFilePath); err != nil {
		e.excelError = _err
		return false
	}
	e.excelCreatedAt = time.Now()

	e.excelFileSizeBytes, _err = file.Size(fullFilePath)
	if _err != nil {
		e.excelError = _err
		return false
	}

	if e.SelfDeleteAfterSeconds != 0 {
		time.AfterFunc(time.Second*time.Duration(e.SelfDeleteAfterSeconds), func() {
			file.Delete(fullFilePath)
		})
	}

	return true
}

func (e *Export) GetExcelFileID() uuid.UUID {
	return e.excelFileID
}

func (e *Export) GetExcelFileName() string {
	return e.excelFileName
}

func (e *Export) GetExcelFullFileName() string {
	return e.excelFullFileName
}

func (e *Export) GetExcelFileExtension() string {
	return e.excelFileExtension
}

func (e *Export) GetExcelFullFilePath() string {
	return e.excelFullFilePath
}

func (e *Export) GetExcelError() error {
	return e.excelError
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
