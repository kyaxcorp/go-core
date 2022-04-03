package record

import (
	"github.com/KyaXTeam/go-core/v2/core/clients/db/scope"
	"gorm.io/gorm"
)

func (r *Record) getDefaultScopes(db *gorm.DB) *gorm.DB {
	// Is Deleted...
	if r.disableDefaultScope.Get() {
		return db
	}

	// Set the default scopes...
	if r.IsStructFieldExists("IsDeleted") {
		db = db.Scopes(scope.IsNotDeleted)
	}

	return db
}
