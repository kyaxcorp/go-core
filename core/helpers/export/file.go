package export

import (
	"github.com/google/uuid"
	"github.com/kyaxcorp/go-core/core/helpers/cmap"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem/tmp"
)

var ExportMap = cmap.New(cmap.MapConstructor{})

func GetExportPath(paths ...string) (string, error) {
	_paths := append([]string{"exporter"}, paths...)
	tmpPath, _err := tmp.GetAppTmpPath(_paths...)
	if _err != nil {
		return "", _err
	}
	return tmpPath, nil
}

func GetExcelFileExportPath() (string, error) {
	tmpPath, _err := GetExportPath("excel")
	if _err != nil {
		return "", _err
	}
	return tmpPath, nil
}

func GetPdfFileExportPath() (string, error) {
	tmpPath, _err := GetExportPath("pdf")
	if _err != nil {
		return "", _err
	}
	return tmpPath, nil
}

func GetHtmlFileExportPath() (string, error) {
	tmpPath, _err := GetExportPath("html")
	if _err != nil {
		return "", _err
	}
	return tmpPath, nil
}

func GetJsonFileExportPath() (string, error) {
	tmpPath, _err := GetExportPath("json")
	if _err != nil {
		return "", _err
	}
	return tmpPath, nil
}

func GetWordFileExportPath() (string, error) {
	tmpPath, _err := GetExportPath("word")
	if _err != nil {
		return "", _err
	}
	return tmpPath, nil
}

/*
ca sa fie usor de exportat, cel mai bine sa generez un obiect cu toate meta datele necesare

*/

type FileExport struct {
	TTL       int64
	Name      string
	Extension string
	//Type string
	// FullPath can be indicated optionally only if you want to overwrite it
	FullPath string
	id       uuid.UUID
}

func NewFileExport(f ...*FileExport) (*FileExport, error) {
	var fe *FileExport
	if len(f) > 0 {
		fe = f[0]
	}

	// Generate a random id for this file export... it will be need when downloading
	id, _err := uuid.NewRandom()
	if _err != nil {
		return nil, _err
	}
	fe.id = id

	return fe, nil
}

func (fe *FileExport) Export() {

	// TODO: continue the work!!! i was lazy!

	//now := time.Now()
	//e.excelFileID = id
	//e.excelFileName = e.ExportName + "_" + conv.Int64ToStr(now.UnixMilli())
	//fullFileName := e.excelFileName + "." + fileExtension

	ExportMap.Set(fe.id.String(), fe.FullPath)
}
