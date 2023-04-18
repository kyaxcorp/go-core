package dbdel

import (
	"context"
	"github.com/google/uuid"
	"github.com/kyaxcorp/go-core/core/clients/db"
	"github.com/kyaxcorp/go-core/core/clients/db/helper"
	"github.com/kyaxcorp/go-core/core/helpers/_struct"
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
)

type DeleteInput struct {
	UserID uuid.UUID
	Ctx    context.Context
	Record interface{}
	DB     *gorm.DB
}

// We created these functions because GORM doesn't allow setting additional data When Calling Delete function!

func DeleteV3(input DeleteInput) (_err error) {
	var dbc *gorm.DB
	if input.DB == nil {
		dbc = db.DBCtx(input.Ctx)
	} else {
		dbc = input.DB
	}

	_struct.SetAny(input.Record, "DeletedByID", &input.UserID)
	_struct.SetAny(input.Record, "DeletedAt", helper.DeletedAtNowP())

	_err = dbc.Transaction(func(tx *gorm.DB) error {
		var txErr error
		if rr, ok := input.Record.(callbacks.BeforeDeleteInterface); ok {
			txErr = rr.BeforeDelete(tx)
			if txErr != nil {
				return txErr
			}
		}
		tx.Statement.SkipHooks = true
		result := tx.Save(input.Record)
		if result.Error != nil {
			return result.Error
		}
		tx.Statement.SkipHooks = false
		if rr, ok := input.Record.(callbacks.AfterDeleteInterface); ok {
			txErr = rr.AfterDelete(tx)
			if txErr != nil {
				return txErr
			}
		}

		return nil
	})

	return
}

func DeleteV2(input DeleteInput) (_err error) {
	var dbc *gorm.DB
	if input.DB == nil {
		dbc = db.DBCtx(input.Ctx)
	} else {
		dbc = input.DB
	}

	_err = dbc.Transaction(func(tx *gorm.DB) error {
		// Update this record!
		tx.Statement.SkipHooks = true
		result := tx.Model(input.Record).Update("deleted_by_id", input.UserID)
		if result.Error != nil {
			return result.Error
		}
		tx.Statement.SkipHooks = false
		result = tx.Delete(input.Record)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	return
}

// Delete -> it will work only in case if UpdateOnSoftDelete
// Tag has been set in Model gorm Tag
func Delete(input DeleteInput) (_err error) {
	var dbc *gorm.DB
	if input.DB == nil {
		dbc = db.DBCtx(input.Ctx)
	} else {
		dbc = input.DB
	}

	_struct.SetAny(input.Record, "DeletedByID", &input.UserID)
	result := dbc.Delete(input.Record)
	if result.Error != nil {
		_err = result.Error
		return
	}
	return
}
