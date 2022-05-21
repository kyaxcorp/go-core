package company

import "github.com/google/uuid"

type FileChunk struct {
	FileID uuid.UUID `gorm:"primaryKey;type:uuid;not null;<-:create;default:gen_random_uuid()"`
	File   File      `gorm:"foreignKey:FileID"`
	// the sequence of the file
	Sequence int64 `gorm:"primaryKey"`
	// Chunk Size Bytes
	Size int64
	//
	Data []byte
}

func (FileChunk) TableName() string {
	return "file_chunks"
}
