package slice

import (
	"github.com/KyaXTeam/go-core/v2/core/helpers/err/define"
	"reflect"
)

func ItemExists(slice interface{}, item interface{}) (bool, error) {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		//panic("Invalid data-type")
		return false, define.Err(0, "invalid data-type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true, nil
		}
	}

	return false, nil
}

func Includes(slice interface{}, item interface{}) (bool, error) {
	return ItemExists(slice, item)
}
