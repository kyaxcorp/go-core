package company

import (
	"github.com/google/uuid"
	"io/fs"
	"time"
)

type File struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;not null;<-:create;default:gen_random_uuid()"`

	// File Name
	Name string
	// File Description
	Description string
	// Size Bytes
	Size string
	// CategoryID -> also an optional field
	CategoryID uuid.UUID
	// Extension -> xls, doc etc...
	Extension string
	// application/json, application/text, etc...
	Type string
	// Physical path of the file on the disk
	Path string
	// If it's related to some Module, it's just a reference
	RelatedToID *uuid.UUID
	RelatedName string
	// Is saved physically or in database, if in database, then it will be stored in another table
	IsPhysical bool

	//EncryptionPassword string
	//EncryptionAlgo     string

	CreatedAt *time.Time `gorm:"type:timestamptz;null;index:idx_core_dates"`
	UpdatedAt *time.Time `gorm:"type:timestamptz;null;index:idx_core_dates"`

	IsDeleted bool       `gorm:"type:bool;not null;default:false;index:idx_core_bool"`
	DeletedAt *time.Time `gorm:"type:timestamptz;null;index:idx_core_dates"`

	CreatedByID *uuid.UUID `gorm:"type:uuid;null"`
	UpdatedByID *uuid.UUID `gorm:"type:uuid;null"`
	DeletedByID *uuid.UUID `gorm:"type:uuid;null"`
}

func (File) TableName() string {
	return "files"
}

/*
We should have functions that:
- Allow us to read and save a file in chunks, the reading will not be performed instantly in the memory, but it will be read
  in chunks... and also save in chunks!
- Allow us to save files from memory
*/

type SaveFileOptions struct {
	// TODO: add defaults?
	ChunkSize int64
}

type SaveFileBytesOptions struct {
	// TODO: add defaults?
	ChunkSize int64
	Name      string
	// TODO: add other file attributes...
}

func SaveFileBytes(data []byte, options ...SaveFileBytesOptions) (*File, error) {
	return nil, nil
}

func SaveFile(file fs.File, options ...SaveFileOptions) (*File, error) {
	var o SaveFileOptions
	if len(options) > 0 {
		o = options[0]
	}
	fileInfo, _err := file.Stat()
	if _err != nil {
		return nil, _err
	}

	data := make([]byte, 5)
	file.Read(data)

	/*
		1. Get a specific file storage location where all files will be saved....
		2. Make a solution that each file will have a very large and mostly infinite ID
		3. create additional primary keys?
		4. think of clustering...
		5.
	*/
}
