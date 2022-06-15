package filter

import (
	"context"
	"gorm.io/gorm"
)

type Counters struct {
	// how many items there are in the db without pagination
	TotalItems int64
	// how many pages there are based on the requested nr of items
	TotalPages int64
	// How many items have been requested
	RequestedNrOfItems int
	// what is the requested page nr
	RequestedPageNr int
	// how many items have been received from the server
	ReceivedNrOfItems int
}

type Details struct {
	// TODO: make a possibility to add more info to the details?!
}

type cachedModel struct {
	// this is the raw structure
	modelName string
	model     interface{}
	// these are the Columns having as index (the model column name) and the value as the GORM Field Name
	dbColumns   map[string]string
	dbTableName string
}

type Input struct {
	PageNr    *int `json:"PageNr"`
	NrOfItems *int `json:"NrOfItems"`
	// maxNrOfItems -> don't allow this param to be controlled by the front end part... it's a security measure!
	maxNrOfItems *int
	// it allows to go higher than 1000 limit!
	Order          []*Order        `json:"Order"`
	Search         *string         `json:"Search"`
	RootConditions *GroupCondition `json:"RootConditions"`

	// here we store the main model name...
	primaryModelName string

	primaryModel cachedModel

	// Here we store the models for knowing the names of the files and to know how to Filter the input...
	models map[string]cachedModel
	// Here we store the map of the models and db fields
	cachedDBFields map[string]DBField

	enableDBFieldsCaching bool

	enableDefaultScope bool
	// This is the database client which after that is used to get the data
	db *gorm.DB
	// dbCounters -> it's the same db, but it's without pagination and data ordering
	dbCounters *gorm.DB
	// getNrOfItems -> defines if it should count the nr of records based on the current filtration statement
	//getNrOfItems   bool
	totalNrOfItems int
	totalNrOfPages int

	ctx context.Context
}

type DBField struct {
	// only field name
	FieldName string
	// field name with table name prefixed
	TableNameFieldName string
}
