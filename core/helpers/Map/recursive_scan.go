package Map

import "reflect"

// it will copy the values from -> to only for existent to keys!!! no more!
func RecursiveCopyToExistent(from map[string]interface{}, to map[string]interface{}) map[string]interface{} {
	for fieldName, fieldValue := range to {
		// Check if from has this field!
		if fromFieldValue, ok := from[fieldName]; ok {
			toMapType := reflect.TypeOf(fieldValue).String()
			fromMapType := reflect.TypeOf(fromFieldValue).String()
			if toMapType == "map[string]interface{}" || toMapType == "map[string]interface {}" {
				// Check if from also has the same field!
				if fromMapType == "map[string]interface{}" || fromMapType == "map[string]interface {}" {
					// it is! copy further!
					to[fieldName] = RecursiveCopyToExistent(fromFieldValue.(map[string]interface{}), fieldValue.(map[string]interface{}))
				}
			} else {
				// it's a value
				to[fieldName] = fromFieldValue
			}
		}
	}
	return to
}
