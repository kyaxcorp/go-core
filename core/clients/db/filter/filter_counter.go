package filter

import (
	"math"
)

/*func (f *Input) CountItems() *Input {
	f.getNrOfItems = true
	return f
}*/

func (f *Input) GetNrOfItems() (*Counters, error) {
	// Query...

	f.check()

	var totalNrOfItems int64

	// Get special db for counters
	db := f.dbCounters.
		Model(f.models[f.primaryModelName].model)

	result := f.applyConditions(db)
	if result != nil {
		db = result
	}
	// It's important to know that GORM removes ORDER BY for counter!
	//db = f.applyOrdering(db)

	if f.enableDefaultScope {
		db = db.Scopes(f.getDefaultScope)
	}

	//Scopes(f.applyConditions).

	// It's important to know that GORM removes ORDER BY for counter!
	result = db.Count(&totalNrOfItems)
	// Let's set the primary model name?!
	//row := f.dbCounters.Select("count(1)").Row()
	if result.Error != nil {
		return nil, result.Error
	}

	totalPages := int64(math.Ceil(float64(totalNrOfItems) / float64(*f.NrOfItems)))

	// Return also Request Page Nr, and Requested Nr Of Items

	// Calculate how many items will be received

	requestedPageNr := *f.PageNr
	requestedNrOfItems := *f.NrOfItems
	// Calculate how many items there are left after setting the pagination and nr of items!
	//leftNrOfItems := int(totalNrOfItems) - (requestedPageNr * requestedNrOfItems)
	// if counted nr of items are lower/equal than requested, and higher than 0, then set the counted one

	endInterval := requestedPageNr * requestedNrOfItems
	startInterval := endInterval - requestedNrOfItems

	receivedNrOfItems := 0
	if startInterval <= int(totalNrOfItems) && int(totalNrOfItems) <= endInterval {
		receivedNrOfItems = int(totalNrOfItems) - startInterval
	} else if int(totalNrOfItems) > endInterval {
		receivedNrOfItems = requestedNrOfItems
	}

	return &Counters{
		TotalItems:         totalNrOfItems,
		TotalPages:         totalPages,
		RequestedNrOfItems: requestedNrOfItems,
		RequestedPageNr:    requestedPageNr,
		ReceivedNrOfItems:  receivedNrOfItems,
	}, nil
}
