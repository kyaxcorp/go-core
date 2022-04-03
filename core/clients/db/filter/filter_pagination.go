package filter

import (
	"github.com/kyaxcorp/go-core/core/clients/db/scope"
	"gorm.io/gorm"
)

func (f *Input) ApplyPagination() *Input {
	f.db = f.db.Scopes(f.applyPagination)
	return f
}

func (f *Input) applyPagination(db *gorm.DB) *gorm.DB {
	db = db.Scopes(
		scope.Paginate(*f.PageNr, *f.NrOfItems),
	)
	return db
}
