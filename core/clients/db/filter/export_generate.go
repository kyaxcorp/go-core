package filter

import (
	//"fmt"
	"github.com/google/uuid"
	"github.com/kyaxcorp/go-core/core/helpers/_struct"
	"github.com/kyaxcorp/go-core/core/helpers/conv"
	"github.com/kyaxcorp/go-core/core/helpers/function"
	"strings"

	//"github.com/kyaxcorp/go-core/core/helpers/err/define"
	"github.com/kyaxcorp/go-core/core/helpers/file"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem/tmp"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"reflect"

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
		Header    string
		FieldName string
		//DBFieldName string
		Handler ExportHandler
		XAxis   string // this is the column
	}

	//firstRow := e.items[0]
	// TODO: later on if no columns defined, we can set to get from the DB or model/table

	var _err error
	//var mFields = make(map[string]string)
	//if e.Model != nil {
	//	mFields, _err = dbHelper.GetModelMapWithDBColumns(e.Model, true)
	//	if _err != nil {
	//		e.excelError = _err
	//		return false
	//	}
	//}

	var headerFields []HeaderField
	for columnNr, columnDetails := range e.Columns {
		clNr := columnNr + 1

		headerName := columnDetails.HeaderName
		//dbColumn := ""
		fieldName := ""

		var _err error
		if columnDetails.FieldName != "" {
			fieldName = columnDetails.FieldName

			// Check if this field exists!

			if !_struct.FieldExists(e.Model, fieldName) {
				log.Println("field doesn't exist!!!", fieldName)
				continue
			}

			//dbColumn, _err = e.Filter.getDBFieldName(columnDetails.FieldName)
			//if _err != nil {
			//	e.excelError = _err
			//	return false
			//}
		} else if function.IsCallable(columnDetails.Handler) {
			// if there is a callback
		} else {
			continue
		}
		//else if columnDetails.DBFieldName != "" {
		//	dbColumn = columnDetails.DBFieldName
		//	if fName, ok := mFields[dbColumn]; ok {
		//		fieldName = fName
		//	} else {
		//		e.excelError = define.Err(0, "structure field not found from database field name -> ", dbColumn)
		//		return false
		//	}
		//}

		if headerName == "" {
			if columnDetails.FieldName != "" {
				headerName = columnDetails.FieldName
			}
			//else {
			//	headerName = columnDetails.DBFieldName
			//}
		}

		xAxis, _err := excelize.ColumnNumberToName(clNr)
		if _err != nil {
			e.excelError = _err
			return false
		}

		headerFields = append(headerFields, HeaderField{
			Header:    headerName,
			FieldName: fieldName,
			//DBFieldName: dbColumn,
			Handler: columnDetails.Handler,
			XAxis:   xAxis,
		})
	}

	for _, headerField := range headerFields {
		excelRowNr := 1

		style, err := f.NewStyle(
			`{
					"alignment":{
						"horizontal":"center"
					},
					"font":{
						"bold":true
					}
				}`,
		)
		if err != nil {
			//fmt.Println(err)
			e.excelError = _err
			return false
		}

		XYAxis := headerField.XAxis + conv.IntToStr(excelRowNr)

		_err := f.SetCellStyle(sheetName, XYAxis, XYAxis, style)
		if _err != nil {
			e.excelError = _err
			return false
		}
		_err = f.SetCellValue(sheetName, XYAxis, headerField.Header)
		if _err != nil {
			e.excelError = _err
			return false
		}
	}

	slice := reflect.ValueOf(e.items)
	sliceLen := slice.Len()

	for rowNr := 0; rowNr < sliceLen; rowNr++ {
		row := slice.Index(rowNr)

		for _, headerField := range headerFields {
			excelRowNr := rowNr + 2
			// row contains the db field names!

			var fieldValue interface{}
			if headerField.FieldName != "" {
				if strings.Contains(headerField.FieldName, ".") {
					//fields := strings.Split(headerField.FieldName, ".")
					fieldValue = _struct.GetNestedFieldReflectValue(row, headerField.FieldName).Interface()
				} else {
					fieldValue = row.FieldByName(headerField.FieldName).Interface()
				}
			} else if headerField.Handler != nil {
				// Execute handler
				fieldValue = headerField.Handler(row)
			} else {
				continue
			}

			//if fieldValue, ok := row[headerField.FieldName]; ok {
			XYAxis := headerField.XAxis + conv.IntToStr(excelRowNr)
			//log.Println("XYAxis", XYAxis, fieldValue)
			_err := f.SetCellValue(sheetName, XYAxis, fieldValue)
			if _err != nil {
				e.excelError = _err
				return false
			}
			//}
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

func (e *Export) GetExcelCreatedAt() time.Time {
	return e.excelCreatedAt
}

func (e *Export) GetExcelFileSizeBytes() int64 {
	return e.excelFileSizeBytes
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
		// TODO: should be deprecated or should be analyzed how to be used...?!...
		e.items = make(map[string]interface{})
		_db = e.Filter.DB().Table(e.TableName)
	} else if e.Model != nil {
		typeOf := reflect.TypeOf(reflect.Indirect(reflect.ValueOf(e.Model)).Interface())
		e.items = reflect.Indirect(reflect.New(reflect.SliceOf(typeOf))).Interface()
		//_db = e.Filter.DB().Model(e.Model)
		_db = e.Filter.DB()
	} else {
		_db = e.Filter.DB()
	}

	// Let's create a slice from the model

	if len(e.Preloads) > 0 {
		for _, preload := range e.Preloads {
			_db = _db.Preload(preload)
		}
	}

	//log.Println("Items2", e.Items2)
	//log.Println("Items2 len",reflect.Indirect(reflect.ValueOf(e.Items2)).Len())
	//log.Println("items", items)
	//log.Println("items len", reflect.Indirect(reflect.ValueOf(items)).Len())

	dbResult := _db.Find(&e.items)
	//log.Println("find finished...")

	if dbResult.Error != nil {
		//log.Println(dbResult.Error.Error())
		return dbResult.Error
	}

	//log.Println(e.items)
	// TODO: we should see what we have here!
	//e.nrOfRows = int64(len(e.items))
	e.nrOfRows = int64(reflect.Indirect(reflect.ValueOf(e.items)).Len())
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

//func (e *Export) GetItems() []map[string]interface{} {
func (e *Export) GetItems() interface{} {
	return e.items
}

//func (e *Export) SetItems(items []map[string]interface{}) *Export {
func (e *Export) SetItems(items interface{}) *Export {
	e.itemsSet = true
	e.items = items
	return e
}
