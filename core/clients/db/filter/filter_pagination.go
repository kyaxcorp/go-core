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
		scope.Paginate(*f.PageNr, *f.NrOfItems, f.OverrideNrOfItemsLimit),
	)
	return db
}

func (f *Input) SetNrOfItems(nrOfItems int, overrideLimit ...bool) *Input {
	if len(overrideLimit) > 0 {
		f.OverrideNrOfItemsLimit = overrideLimit[0]
	}
	f.NrOfItems = &nrOfItems
	return f
}

func (f *Input) SetPageNr(pageNr int) *Input {
	f.PageNr = &pageNr
	return f
}
