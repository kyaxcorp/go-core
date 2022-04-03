package filter

import (
	"github.com/KyaXTeam/go-core/v2/core/clients/db/scope"
	"gorm.io/gorm"
)

func (f *Input) EnableDefaultScope() *Input {
	f.enableDefaultScope = true
	return f
}

func (f *Input) DisableDefaultScope() *Input {
	f.enableDefaultScope = false
	return f
}

func (f *Input) applyDefaultScope() *Input {
	f.db = f.db.Scopes(f.getDefaultScope)
	return f
}

func (f *Input) getDefaultScope(db *gorm.DB) *gorm.DB {
	// Check if field exists! (check all models?!)
	// We should check the model?!
	// TODO: should we cache here anything?!

	if _, _err := f.getDBFieldName("IsDeleted"); _err == nil {
		db.Scopes(scope.IsNotDeleted)
	}

	return db
}
