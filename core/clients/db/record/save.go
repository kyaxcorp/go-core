package record

import (
	"github.com/google/uuid"
	"github.com/kyaxcorp/go-core/core/helpers/Map"
	"github.com/kyaxcorp/go-core/core/helpers/_interface"
	"github.com/kyaxcorp/go-core/core/helpers/_struct"
	"github.com/kyaxcorp/go-core/core/helpers/err/define"
	"github.com/kyaxcorp/go-core/core/helpers/json"
	"gorm.io/gorm"
	"reflect"
	"time"
)

/*
Create some kind of structure which will be embedded in another...?!
It will take the functions, but it will not take the values, because it will not know the parent structure!
Only if we create a function which sets the reference to the parent... in this case it may work...

*/

// TODO: we can set global hooks for this structure...

func (r *Record) getOmitFields() []string {
	_strMap := _struct.New(r.modelStruct).Map()
	var omitFields []string

	// TODO: should we also omit for nested structures that have not been set?!
	// TODO: or we should simply load their data and add to it...

	// gorm allows to omit sub fields! they should be declared with . (dots)
	for fieldName, _ := range _strMap {
		if _, fieldFound := r.saveData[fieldName]; !fieldFound {
			// if something is not present, then omit it
			omitFields = append(omitFields, fieldName)
		}
	}
	return omitFields
}

func (r *Record) GetSaveData() map[string]interface{} {
	return r.saveData
}

// TODO: get save data field and set save data field!

func (r *Record) SetSaveData(saveData map[string]interface{}) *Record {
	r.saveData = saveData
	return r
}

func (r *Record) generateSaveDataModel() interface{} {
	_model := _interface.CloneInterfaceItem(r.modelStruct)
	_json, _err := json.Encode(r.saveData)
	if _err != nil {
		panic("failed to encode r.saveData to Json -> " + _err.Error())
	}
	_err = json.Decode(_json, _model)
	if _err != nil {
		panic("failed to Decode _json to _model -> " + _err.Error())
	}
	return _model
}

func (r *Record) Save() bool {
	_db := r.getDB()

	var result *gorm.DB

	if !r.prepareSaveData() {
		return false
	}

	r.callOnBeforeSave()

	uID := r.GetUserID()
	uIDisNil := r.isUserIDNil()

	// TODO: la noi ID e uuid.UUID insa in input el vine ca string

	// Updated should be always present!
	if !uIDisNil && _struct.FieldExists(r.modelStruct, "UpdatedBy") {
		r.saveData["UpdatedBy"] = uID
	}
	//
	if _struct.FieldExists(r.modelStruct, "UpdatedAt") {
		r.saveData["UpdatedAt"] = time.Now()
	}

	if r.IsCreateMode() {
		// If it's nil, then we should create it!
		if !uIDisNil && _struct.FieldExists(r.modelStruct, "CreatedBy") {
			// check which type is user id -> uuid or other type
			r.saveData["CreatedBy"] = uID
		}
		//
		if _struct.FieldExists(r.modelStruct, "CreatedAt") {
			r.saveData["CreatedAt"] = time.Now()
		}

		// 1. copy the data to the real structure
		// 2. omit all fields that are not in the list (we should not include primary keys in the omit list!) because we should receive them back!
		// 3. Create
		// 4. Now let's read the data only by list and set it to dbData!
		// 5. launch reload to load the entire data!

		saveDataModel := r.generateSaveDataModel()
		result = _db.Omit(r.getOmitFields()...).Create(saveDataModel)
		r.dbData = saveDataModel
		// TODO: later on we should do a reload like on save?!

		r.callOnAfterInsert()
	} else {
		r.callOnBeforeUpdate()
		saveDataModel := r.generateSaveDataModel()
		result = _db.Omit(r.getOmitFields()...).Save(saveDataModel)
		r.dbData = saveDataModel
		r.ReloadData()

		// We should update it!
		//result = _db.Save(r.saveData)
		r.callOnAfterUpdate()
	}
	r.callOnAfterSave()

	// We can return BOOL and save the error somewhere in the record as the last error!

	if result.Error != nil {
		r.setDBError(result.Error)
		r.callOnError()
		r.callOnDBError()
		r.callOnSaveError()
		return false
	}

	// Load back the data!

	return true
}

func (r *Record) isUserIDNil() bool {
	uID := r.GetUserID()
	uIDv := reflect.ValueOf(uID)
	uIDType := uIDv.Type().String()
	uIDisNil := false
	if uIDType == "uuid.UUID" || uIDType == "*uuid.UUID" {
		// TODO: we should test if it's nil by using pointer or not
		if uID == uuid.Nil {
			uIDisNil = true
		}
	} else {
		if uID == nil {
			uIDisNil = true
		}
	}
	return uIDisNil
}

func (r *Record) GetSavedData() interface{} {
	//return r.saveData
	return r.dbData
}

func (r *Record) prepareSaveData() bool {
	callSimpleError := func(_err error) {
		r.setError(_err)
		r.callOnError()
		r.callOnSaveError()
	}

	// let's recreate the map with no keys...
	r.saveData = make(map[string]interface{})

	if r.IsSaveMode() {
		// if it's save mode then we should get the loaded data from r.dbData
		// and after that put over it the inputData

		dbDataMap := _struct.New(r.dbData).Map()
		//r.dataMap
		// copy first the current db data to saveData
		Map.CopyStringInterface(dbDataMap, r.saveData)
	}

	// let's copy the inputData to save data, why? because we don't want to flood the inputData with other information
	// the saveData variable can have or can be supplied with other additional information!
	//Map.CopyStringInterface(r.inputData, r.saveData)
	Map.CopyStringInterface(r.dataMap, r.saveData)

	// step 3 - copy the data from the input
	switch r.inputDataType {
	case inputDataMapInterface:

	case inputDataStruct:

	default:
		callSimpleError(define.Err(0, "unknown input data"))
		return false
	}
	// Check some fields if the types are correct....
	return true
}
