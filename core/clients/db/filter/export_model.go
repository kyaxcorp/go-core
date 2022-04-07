package filter

type Export struct {
	filter *Input
	// items -> here we will receive the queried items
	items []map[string]interface{}
	// if items are set
	itemsSet bool

	nrOfRows    int64
	nrOfColumns int64

	columns []string

	// Excel part
	excelFileName      string
	excelFullFileName  string
	excelFileExtension string
	excelFileSizeBytes int64
	excelError         error

	// TODO....
	pdfFileName string
}
