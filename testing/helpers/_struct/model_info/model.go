package main

import "github.com/google/uuid"

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
}

type Test struct {
	Model       interface{}
	ModelStruct interface{}
}
