package slice

import (
	"github.com/kyaxcorp/go-core/core/helpers/err/define"
	"reflect"
)

func IndexExists(slice interface{}, indexNr int) (bool, error) {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		// panic("Invalid data-type")
		return false, define.Err(0, "invalid data-type")
	}

	return s.Len() > indexNr, nil
}
