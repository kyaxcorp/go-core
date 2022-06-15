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
		scope.Paginate(*f.PageNr, *f.NrOfItems, *f.maxNrOfItems),
	)
	return db
}

func (f *Input) SetNrOfItems(nrOfItems int) *Input {
	f.NrOfItems = &nrOfItems
	return f
}

func (f *Input) SetMaxNrOfItems(maxNrOfItems int) *Input {
	f.maxNrOfItems = &maxNrOfItems
	return f
}

func (f *Input) SetPageNr(pageNr int) *Input {
	f.PageNr = &pageNr
	return f
}
