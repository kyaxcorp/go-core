package _struct

import (
	"fmt"
	"reflect"
)

// FillMap is the same as Map. Instead of returning the output, it fills the
// given map.
func (h *Helper) FillMap(out map[string]interface{}) {
	if out == nil {
		return
	}

	fields := h.structFields()

	for _, field := range fields {
		name := field.Name
		val := h.valueOf.FieldByName(name)
		isSubStruct := false
		var finalVal interface{}

		tagName, tagOpts := parseTag(field.Tag.Get(h.TagName))
		if tagName != "" {
			name = tagName
		}

		// if the value is a zero value and the field is marked as omitempty do
		// not include
		if tagOpts.Has("omitempty") {
			zero := reflect.Zero(val.Type()).Interface()
			current := val.Interface()

			if reflect.DeepEqual(current, zero) {
				continue
			}
		}

		if !tagOpts.Has("omitnested") {
			finalVal = h.nested(val)

			v := reflect.ValueOf(val.Interface())
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}

			switch v.Kind() {
			case reflect.Map, reflect.Struct:
				isSubStruct = true
			}
		} else {
			finalVal = val.Interface()
		}

		if tagOpts.Has("string") {
			s, ok := val.Interface().(fmt.Stringer)
			if ok {
				out[name] = s.String()
			}
			continue
		}

		if isSubStruct && (tagOpts.Has("flatten")) {
			for k := range finalVal.(map[string]interface{}) {
				out[k] = finalVal.(map[string]interface{})[k]
			}
		} else {
			out[name] = finalVal
		}
	}
}

// Map converts the given struct to a map[string]interface{}, where the keys
// of the map are the field names and the values of the map the associated
// values of the fields. The default key string is the struct field name but
// can be changed in the struct field's tag value. The "structs" key in the
// struct's field tag value is the key name. Example:
//
//   // Field appears in map as key "myName".
//   Name string `structs:"myName"`
//
// A tag value with the content of "-" ignores that particular field. Example:
//
//   // Field is ignored by this package.
//   Field bool `structs:"-"`
//
// A tag value with the content of "string" uses the stringer to get the value. Example:
//
//   // The value will be output of Animal's String() func.
//   // Map will panic if Animal does not implement String().
//   Field *Animal `structs:"field,string"`
//
// A tag value with the option of "flatten" used in a struct field is to flatten its fields
// in the output map. Example:
//
//   // The FieldStruct's fields will be flattened into the output map.
//   FieldStruct time.Time `structs:",flatten"`
//
// A tag value with the option of "omitnested" stops iterating further if the type
// is a struct. Example:
//
//   // Field is not processed further by this package.
//   Field time.Time     `structs:"myName,omitnested"`
//   Field *http.Request `structs:",omitnested"`
//
// A tag value with the option of "omitempty" ignores that particular field if
// the field value is empty. Example:
//
//   // Field appears in map as key "myName", but the field is
//   // skipped if empty.
//   Field string `structs:"myName,omitempty"`
//
//   // Field appears in map as key "Field" (the default), but
//   // the field is skipped if empty.
//   Field string `structs:",omitempty"`
//
// Note that only exported fields of a struct can be accessed, non exported
// fields will be neglected.
func (h *Helper) Map() map[string]interface{} {
	out := make(map[string]interface{})
	h.FillMap(out)
	return out
}
