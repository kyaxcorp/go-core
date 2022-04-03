package main

import (
	"github.com/google/uuid"
	"github.com/kyaxcorp/go-core/core/helpers/_struct"
	"log"
)

type Terminal struct {
	// PRIMARY KEYS ARE OK!
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;not null;<-:create;default:gen_random_uuid()"`
	PosID     uuid.UUID `gorm:"primaryKey;type:uuid;not null;<-:create"`
	CompanyID uuid.UUID `gorm:"primaryKey;type:uuid;not null;<-:create"`

	// Name -> a simple naming offered for this device...
	Name string `gorm:"size:200;type:string;null"`
	// Some info about this terminal
	Description string `gorm:"size:1000;type:string;null"`

	// Type : Virtual, Physical
	Type uint8
	// IsStationary or it's mobile... (meaning that is movable)
	IsStationary bool

	CreatedBy uuid.UUID
}

type Details struct {
	UserID interface{}
}

func main() {

	_id, _ := uuid.NewRandom()
	d := &Details{
		UserID: _id,
	}

	t := &Terminal{}
	tStr := Terminal{}

	if _struct.FieldExists(tStr, "CreatedBy") && d.UserID != nil && _struct.New(t).SetAny("CreatedBy", d.UserID) {
		log.Println("success", t)
	} else {
		log.Println("failed", t)
	}

}
