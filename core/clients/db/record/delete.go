package record

import (
	"github.com/kyaxcorp/go-core/core/helpers/_struct"
	"github.com/kyaxcorp/go-core/core/helpers/err"
	"gorm.io/gorm"
	"time"
)

func (r *Record) Delete() bool {
	_db := r.getDB()

	if !r.prepareSaveData() {
		return false
	}

	uID := r.GetUserID()
	uIDisNil := r.isUserIDNil()

	r.callOnBeforeDelete()
	var result *gorm.DB
	if r.IsCreateMode() {
		_err := err.New(0, "record doesn't exist or id not set")
		r.setDBError(_err)
		return false
	} else {
		// We should update it!
		// TODO: should we get the record before making changes?!...
		if _struct.FieldExists(r.modelStruct, "IsDeleted") {
			//r.saveData["IsDeleted"] = true
			_struct.New(r.saveData).SetInterface("IsDeleted", true)

			if _struct.FieldExists(r.modelStruct, "DeletedAt") {
				//r.saveData["DeletedAt"] = time.Now()
				_struct.New(r.saveData).SetInterface("DeletedAt", time.Now())
			}
			if !uIDisNil && _struct.FieldExists(r.modelStruct, "DeletedBy") {
				//r.saveData["DeletedBy"] = uID
				_struct.New(r.saveData).SetInterface("DeletedBy", uID)
			}

			r.callOnBeforeUpdate()
			//saveDataModel := r.generateSaveDataModel()
			result = _db.Omit(r.getOmitFields()...).Save(r.saveData)
			r.loadDataForUpdate = false
			r.dbData = r.saveData

			//r.ReloadData()
			r.callOnAfterUpdate()

			if result.Error != nil {
				r.setDBError(result.Error)
				return false
			}
			r.callOnAfterDelete()
			return true
		} else {
			return r.ForceDelete()
		}
	}
}

// ForceDelete -> it will completely delete the record from the DB
func (r *Record) ForceDelete() bool {
	_db := r.getDB()
	// TODO: should we get the record before making changes?!...
	// TODO: should we check if exists in the DB?!...or gorm does it...

	r.callOnBeforeForceDelete()

	var result *gorm.DB
	if r.IsCreateMode() {
		_err := err.New(0, "record doesn't exist or id not set")
		r.setDBError(_err)
		r.callOnError()
		r.callOnDeleteError()
		return false
	} else {
		// We should delete it!
		// Does gorm know the primary keys?! when there are multiple
		result = _db.Delete(&r.saveData)
		if result.Error != nil {
			r.setDBError(result.Error)
			r.callOnError()
			r.callOnDBError()
			r.callOnDeleteError()
			return false
		}
		r.callOnAfterForceDelete()
		return true
	}
	// TODO: add hooks
}
