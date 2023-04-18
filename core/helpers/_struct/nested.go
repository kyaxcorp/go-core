package _struct

import (
	"github.com/kyaxcorp/go-core/core/helpers/_int"
	"github.com/kyaxcorp/go-core/core/helpers/errors2/define"
	"reflect"
	"strings"
)

func GetNestedFieldReflectValue(val reflect.Value, path string) (reflect.Value, error) {
	if path == "" {
		//panic("path is empty")
		return reflect.Value{}, define.Err(0, "path is empty")
	}
	keys := strings.Split(path, ".")
	nrOfKeys := len(keys)

	//log.Println("nr of keys", nrOfKeys)
	//log.Println("keys", keys)

	if nrOfKeys == 0 {
		//panic("nr of keys is 0")
		return reflect.Value{}, define.Err(0, "nr of keys is 0")
	}

	i := reflect.Indirect(val)
	if i == (reflect.Value{}) {
		//log.Println("IT'S ZERO VALUE")
		return reflect.Value{}, define.Err(0, "i is zero value")
	}

	firstKey := keys[0]
	nrOfRemainingKeys := nrOfKeys - 1
	//log.Println("nr of remaining keys", nrOfRemainingKeys)
	//log.Println("first key", firstKey)

	// Check if key is a number, string
	// check if string is key for map... maybe we should write this in other way!

	if _int.IsNumber(firstKey) {
		// TODO:
		return reflect.Value{}, nil
	} else {

		objVal := i.FieldByName(firstKey)
		if objVal == (reflect.Value{}) {
			panic("field " + firstKey + " doesn't exist")
		}

		if nrOfRemainingKeys <= 0 {
			// it's the last one!
			return objVal, nil
		} else {
			var remainKeys = make([]string, nrOfRemainingKeys)
			copy(remainKeys[0:nrOfRemainingKeys], keys[1:nrOfRemainingKeys+1])
			newPath := strings.Join(remainKeys, ".")
			//log.Println("copy these:", keys[1:nrOfRemainingKeys])
			//log.Println("copy these:", keys[1:nrOfRemainingKeys+1])
			//log.Println("remain keys", remainKeys)
			return GetNestedFieldReflectValue(objVal, newPath)
		}
	}
}

// nested retrieves recursively all types for the given value and returns the
// nested value.
func (h *Helper) nested(val reflect.Value) interface{} {
	var finalVal interface{}

	v := reflect.ValueOf(val.Interface())
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		n := New(val.Interface())
		n.TagName = h.TagName
		m := n.Map()

		// do not add the converted value if there are no exported fields, ie:
		// time.Time
		if len(m) == 0 {
			finalVal = val.Interface()
		} else {
			finalVal = m
		}
	case reflect.Map:
		// get the element type of the map
		mapElem := val.Type()
		switch val.Type().Kind() {
		case reflect.Ptr, reflect.Array, reflect.Map,
			reflect.Slice, reflect.Chan:
			mapElem = val.Type().Elem()
			if mapElem.Kind() == reflect.Ptr {
				mapElem = mapElem.Elem()
			}
		}

		// only iterate over struct types, ie: map[string]StructType,
		// map[string][]StructType,
		if mapElem.Kind() == reflect.Struct ||
			(mapElem.Kind() == reflect.Slice &&
				mapElem.Elem().Kind() == reflect.Struct) {
			m := make(map[string]interface{}, val.Len())
			for _, k := range val.MapKeys() {
				m[k.String()] = h.nested(val.MapIndex(k))
			}
			finalVal = m
			break
		}

		// TODO(arslan): should this be optional?
		finalVal = val.Interface()
	case reflect.Slice, reflect.Array:
		if val.Type().Kind() == reflect.Interface {
			finalVal = val.Interface()
			break
		}

		// TODO(arslan): should this be optional?
		// do not iterate of non struct types, just pass the value. Ie: []int,
		// []string, co... We only iterate further if it's a struct.
		// i.e []foo or []*foo
		if val.Type().Elem().Kind() != reflect.Struct &&
			!(val.Type().Elem().Kind() == reflect.Ptr &&
				val.Type().Elem().Elem().Kind() == reflect.Struct) {
			finalVal = val.Interface()
			break
		}

		slices := make([]interface{}, val.Len())
		for x := 0; x < val.Len(); x++ {
			slices[x] = h.nested(val.Index(x))
		}
		finalVal = slices
	default:
		finalVal = val.Interface()
	}

	return finalVal
}
