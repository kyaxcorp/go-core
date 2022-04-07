package filter

import (
	"github.com/google/uuid"
	"github.com/kyaxcorp/go-core/core/helpers/conv"
	"github.com/kyaxcorp/go-core/core/helpers/file"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem/tmp"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"log"
	"time"
)

func (e *Export) GeneratePdf() {

}

func (e *Export) GetExportPath(paths ...string) (string, error) {
	_paths := append([]string{"exporter"}, paths...)
	tmpPath, _err := tmp.GetAppTmpPath(_paths...)
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

	if e.nrOfRows == 0 {
		return false
	}

	// Let's initially export everything...
	f := excelize.NewFile()

	sheetName := "Sheet1"
	sheetIndex := f.GetSheetIndex(sheetName)

	// First, let's create the struct for speeding up the process!

	type HeaderField struct {
		Header      string
		DBFieldName string
		XAxis       string // this is the column
	}

	//firstRow := e.items[0]
	// TODO: later on if no columns defined, we can set to get from the DB or model/table

	var headerFields []HeaderField
	for columnNr, columnDetails := range e.Columns {
		headerName := columnDetails.HeaderName
		dbColumn := ""
		var _err error
		if columnDetails.FieldName != "" {
			dbColumn, _err = e.Filter.getDBFieldName(columnDetails.FieldName)
			if _err != nil {
				e.excelError = _err
				return false
			}
		} else if columnDetails.DBFieldName != "" {
			dbColumn = columnDetails.DBFieldName
		} else {
			continue
		}

		if headerName == "" {
			if columnDetails.FieldName != "" {
				headerName = columnDetails.FieldName
			} else {
				headerName = columnDetails.DBFieldName
			}
		}

		xAxis, _err := excelize.ColumnNumberToName(columnNr)
		if _err != nil {
			e.excelError = _err
			return false
		}

		headerFields = append(headerFields, HeaderField{
			Header:      headerName,
			DBFieldName: dbColumn,
			XAxis:       xAxis,
		})
	}

	for rowNr, row := range e.items {
		for _, headerField := range headerFields {
			// row contains the db field names!
			if fieldValue, ok := row[headerField.DBFieldName]; ok {
				XYAxis := headerField.XAxis + conv.IntToStr(rowNr)
				log.Println("XYAxis", XYAxis, fieldValue)
				_err := f.SetCellValue(sheetName, XYAxis, fieldValue)
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
	//fileName := id.String()

	// TODO: we should save the file id into a memory stack with id's... maybe in the filter somewhere...

	if e.ExportName == "" {
		e.ExportName = id.String()
	}

	now := time.Now()
	e.excelFileID = id
	e.excelFileName = e.ExportName + "_" + conv.Int64ToStr(now.UnixMilli())
	fullFileName := e.excelFileName + "." + fileExtension

	e.excelFullFileName = fullFileName
	e.excelFileExtension = fileExtension

	tmpPath, _err := e.GetExcelExportPath()
	if _err != nil {
		e.excelError = _err
		return false
	}

	fullFilePath := tmpPath + filesystem.DirSeparator() + fullFileName
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
			// Check if exists...!?
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

	var _db *gorm.DB
	if e.TableName != "" {
		_db = e.Filter.DB().Table(e.TableName)
	} else if e.Model != nil {
		_db = e.Filter.DB().Model(e.Model)
	} else {
		_db = e.Filter.DB()
	}

	dbResult := _db.Find(&e.items)

	if dbResult.Error != nil {
		return dbResult.Error
	}
	e.nrOfRows = int64(len(e.items))
	e.itemsSet = true
	return nil
}

func (e *Export) SetColumns(columns []ExportColumn) *Export {
	e.Columns = columns
	return e
}

func (e *Export) GetNrOfRows() int64 {
	return e.nrOfRows
}

func (e *Export) GetItems() []map[string]interface{} {
	return e.items
}

func (e *Export) SetItems(items []map[string]interface{}) *Export {
	e.itemsSet = true
	e.items = items
	return e
}
