package scope

import (
	"gorm.io/gorm"
)

func Paginate(pageNr int, nrOfItems int, overrideLimit ...bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		_nrOfItems := nrOfItems
		switch {
		case _nrOfItems > 1000:
			// Don't allow more than 1000 items!
			_nrOfItems = 1000
			if len(overrideLimit) > 0 {
				if overrideLimit[0] {
					_nrOfItems = nrOfItems
				}
			}
		case _nrOfItems <= 0:
			// If it's lower or equal than 0, then give only 10 items as the default value
			_nrOfItems = 10
		}

		offset := (pageNr - 1) * _nrOfItems
		return db.Offset(offset).Limit(_nrOfItems)
	}
}
