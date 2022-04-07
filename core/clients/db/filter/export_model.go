package filter

import (
	"github.com/google/uuid"
	"time"
)

type Export struct {
	// This is a prefix for exporting/generating the file!
	ExportName string

	Filter *Input
	// It will delete itself after a period of time...
	// if 0, then it will not be deleted!
	SelfDeleteAfterSeconds int64
	// items -> here we will receive the queried items
	items []map[string]interface{}
	// if items are set
	itemsSet bool

	nrOfRows    int64
	nrOfColumns int64

	columns []string

	// Excel part
	excelFileID        uuid.UUID
	excelFileName      string
	excelFullFileName  string
	excelFullFilePath  string
	excelFileExtension string
	excelCreatedAt     time.Time
	excelFileSizeBytes int64
	excelError         error

	// TODO....
	pdfFileName string
}
