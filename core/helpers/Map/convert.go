package Map

import (
	"database/sql"
	"database/sql/driver"
	"github.com/KyaXTeam/go-core/v2/core/helpers/_struct"
	"reflect"
)

type MapConvertOpts struct {
	IfFailedConversionSetValFromStruct bool
	DeepCopy                           bool
}

func ConvertMapValuesBasedOnModel(
	input map[string]interface{},
	modelStruct interface{},
	opts *MapConvertOpts,
) (map[string]interface{}, error) {
	m := _struct.New(modelStruct)

	deepCopy := true
	ifFailedConversionSetValFromStruct := true
	if opts != nil {
		ifFailedConversionSetValFromStruct = opts.IfFailedConversionSetValFromStruct
		deepCopy = opts.DeepCopy
	}

	rm := make(map[string]interface{})

	for fieldName, fieldVal := range input {
		t := m.GetFieldReflectType(fieldName)
		t, isPtr := indirectType(t)
		//log.Println("field name", fieldName, " -> type", t, t.Name(), " -> is ptr -> ", isPtr)

		var newVal reflect.Value
		//newValRef := reflect.New(t)
		//log.Println("field type -> ", t.Name())
		//newVal := newValRef.Interface()
		//log.Print(newVal)

		valRef := reflect.ValueOf(fieldVal)
		valType := reflect.TypeOf(fieldVal)

		//log.Println("field name", fieldName, valType)

		if valRef == (reflect.Value{}) {
			// is nil, then we should make the value also NIL!

			newVal = indirect(reflect.New(t))

			// set the default value of the struct...? but what if it's nil!?
			if isPtr {
				rm[fieldName] = nil
			} else {
				rm[fieldName] = newVal.Interface()
			}
			continue
		}

		// Check if not same type!
		if t != valType {
			if valRef.CanConvert(t) {
				//log.Println("converting ", fieldName)
				newVal = valRef.Convert(t)
				rm[fieldName] = newVal.Interface()
			} else {
				//log.Println("cant convert ", fieldName)
				newVal = indirect(reflect.New(t))

				//log.Println("new val", fieldName, newVal)

				//copier.Copy()
				if set(newVal, valRef, deepCopy) {
					//log.Println("copy success", fieldName)
					rm[fieldName] = newVal.Interface()
				} else {

					//log.Println("copy failed", fieldName)
					if ifFailedConversionSetValFromStruct {
						rm[fieldName] = newVal.Interface()
					}
				}
			}
		} else {
			//log.Print("conversion unnecessary!", fieldName)
			rm[fieldName] = fieldVal
		}
	}
	return rm, nil
}

func driverValuer(v reflect.Value) (i driver.Valuer, ok bool) {

	if !v.CanAddr() {
		i, ok = v.Interface().(driver.Valuer)
		return
	}

	i, ok = v.Addr().Interface().(driver.Valuer)
	return
}

func indirectType(reflectType reflect.Type) (_ reflect.Type, isPtr bool) {
	for reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Slice {
		reflectType = reflectType.Elem()
		isPtr = true
	}
	return reflectType, isPtr
}

func indirect(reflectValue reflect.Value) reflect.Value {
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue
}

// TODO: this function should be made as standard here in the lib!
func set(to, from reflect.Value, deepCopy bool) bool {
	if from.IsValid() {

		if to.Kind() == reflect.Ptr {
			// set `to` to nil if from is nil
			if from.Kind() == reflect.Ptr && from.IsNil() {
				to.Set(reflect.Zero(to.Type()))
				return true
			} else if to.IsNil() {
				// `from`         -> `to`
				// sql.NullString -> *string
				if fromValuer, ok := driverValuer(from); ok {
					v, err := fromValuer.Value()
					if err != nil {
						return false
					}
					// if `from` is not valid do nothing with `to`
					if v == nil {
						return true
					}
				}
				// allocate new `to` variable with default value (eg. *string -> new(string))
				to.Set(reflect.New(to.Type().Elem()))
			}
			// depointer `to`
			to = to.Elem()
		}

		if deepCopy {
			toKind := to.Kind()
			if toKind == reflect.Interface && to.IsNil() {
				if reflect.TypeOf(from.Interface()) != nil {
					to.Set(reflect.New(reflect.TypeOf(from.Interface())).Elem())
					toKind = reflect.TypeOf(to.Interface()).Kind()
				}
			}
			if from.Kind() == reflect.Ptr && from.IsNil() {
				return true
			}
			if toKind == reflect.Struct || toKind == reflect.Map || toKind == reflect.Slice {
				return false
			}
		}

		if from.Type().ConvertibleTo(to.Type()) {
			to.Set(from.Convert(to.Type()))
		} else if toScanner, ok := to.Addr().Interface().(sql.Scanner); ok {
			// `from`  -> `to`
			// *string -> sql.NullString
			if from.Kind() == reflect.Ptr {
				// if `from` is nil do nothing with `to`
				if from.IsNil() {
					return true
				}
				// depointer `from`
				from = indirect(from)
			}
			// `from` -> `to`
			// string -> sql.NullString
			// set `to` by invoking method Scan(`from`)
			err := toScanner.Scan(from.Interface())
			if err != nil {
				return false
			}
		} else if fromValuer, ok := driverValuer(from); ok {
			// `from`         -> `to`
			// sql.NullString -> string
			v, err := fromValuer.Value()
			if err != nil {
				return false
			}
			// if `from` is not valid do nothing with `to`
			if v == nil {
				return true
			}
			rv := reflect.ValueOf(v)
			if rv.Type().AssignableTo(to.Type()) {
				to.Set(rv)
			}
		} else if from.Kind() == reflect.Ptr {
			return set(to, from.Elem(), deepCopy)
		} else {
			return false
		}
	}

	return true
}
