package scope

import (
	"gorm.io/gorm"
)

func Paginate(pageNr int, nrOfItems int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		switch {
		case nrOfItems > 1000:
			// Don't allow more than 1000 items!
			nrOfItems = 1000
		case nrOfItems <= 0:
			// If it's lower or equal than 0, then give only 10 items as the default value
			nrOfItems = 10
		}

		offset := (pageNr - 1) * nrOfItems
		return db.Offset(offset).Limit(nrOfItems)
	}
}
