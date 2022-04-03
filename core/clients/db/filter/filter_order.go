package filter

import (
	"gorm.io/gorm"
	"strings"
)

func (f *Input) ApplyOrdering() *Input {
	f.db = f.db.Scopes(f.applyOrdering)
	return f
}

func (f *Input) applyOrdering(db *gorm.DB) *gorm.DB {
	// Loop through ordering...

	if f.Order == nil || len(f.Order) == 0 {
		return db
	}

	for _, o := range f.Order {
		// Check if field exists?!

		// Checking against injection!
		ord := o.FieldName
		if !validateFieldName(ord) {
			panic("invalid order field name format -> " + ord)
		}
		// Transform the field name
		transformedOrd := f.getDBFieldNameOrPanic(ord)

		// Direction may be empty!
		if o.Direction != nil && *o.Direction != "" {
			direction := strings.ToLower(*o.Direction)
			switch direction {
			case "asc":
			case "desc":
			default:
				panic("invalid order direction format -> " + direction)
			}

			transformedOrd += " " + direction
		}
		db = db.Order(transformedOrd)
	}

	// TODO: maybe set a default order...
	//f.db.Scopes(
	//	scope.OrderByCreatedAtDesc,
	//)
	return db
}
