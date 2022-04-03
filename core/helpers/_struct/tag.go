package _struct

import (
	"strings"
)

const (
	InvalidObjType        = "obj type should be non-pointer and type struct"
	InvalidPointerObjType = "obj type should be pointer and type struct"
)

type InputRef struct {
	// you can choose what to set, one of these options
	//Helper *Helper
	TagVal string
}

func (h *Helper) GetTagValByInputRef(inputRef *InputRef, fieldName string, tagName string) string {
	if inputRef == nil {
		return h.GetFieldTagValue(fieldName, tagName)
	}

	var tagVal string
	if inputRef.TagVal != "" {
		tagVal = inputRef.TagVal
	}
	return tagVal
}

func (h *Helper) GetFieldTagKeyValue(fieldName string, tagName string, tagKey string) string {
	tagVal := h.GetTagValByInputRef(nil, fieldName, tagName)

	// Check if the key even exists...
	if !strings.Contains(tagVal, tagKey) {
		return ""
	}

	keys := strings.Split(tagVal, ";")
	if len(keys) == 0 {
		return ""
	}
	for _, v := range keys {
		keySplit := strings.Split(v, ":")
		keySplitLen := len(keySplit)
		if keySplitLen == 0 {
			continue
		}
		if keySplit[0] == tagKey {
			if keySplitLen == 2 {
				return keySplit[1]
			}
		}
	}

	return ""
}

func GetFieldTagKeyValue(obj interface{}, fieldName string, tagName string, tagKey string) string {
	return New(obj).GetFieldTagKeyValue(fieldName, tagName, tagKey)
}

func (h *Helper) GetFieldTagValue(fieldName string, tagName string) string {
	// TODO: we should revise that! we should allow with pointer or plain!
	if !IsPlainStruct(h.NonPtrObj) {
		panic(InvalidObjType)
	}
	f, isFound := h.typeOf.FieldByName(fieldName)
	if !isFound {
		return ""
	}
	return f.Tag.Get(tagName)
}

func (h *Helper) GetFieldTag(fieldName string) string {
	if !IsPlainStruct(h.NonPtrObj) {
		panic(InvalidObjType)
	}
	f, isFound := h.typeOf.FieldByName(fieldName)
	if !isFound {
		return ""
	}
	return string(f.Tag)
}

func (h *Helper) IsFieldTagExistsExt(fieldName string, tagName string) (tagExists bool, isFieldFound bool) {
	if !IsPlainStruct(h.NonPtrObj) {
		panic(InvalidObjType)
	}
	f, _isFound := h.typeOf.FieldByName(fieldName)
	if !_isFound {
		return false, false
	}
	_, isFound := f.Tag.Lookup(tagName)
	return isFound, true
}

// IsFieldTagExistsExt -> extended
func IsFieldTagExistsExt(obj interface{}, fieldName string, tagName string) (tagExists bool, isFieldFound bool) {
	return New(obj).IsFieldTagExistsExt(fieldName, tagName)
}

// IsFieldTagExists -> standard version
func IsFieldTagExists(obj interface{}, fieldName string, tagName string) bool {
	return New(obj).IsFieldTagExists(fieldName, tagName)
}

func (h *Helper) IsFieldTagExists(fieldName string, tagName string) bool {
	if !IsPlainStruct(h.NonPtrObj) {
		panic(InvalidObjType)
	}
	f, _isFound := h.typeOf.FieldByName(fieldName)
	if !_isFound {
		return false
	}
	_, isFound := f.Tag.Lookup(tagName)
	return isFound
}

func (h *Helper) IsFieldTagKeyExists(inputRef *InputRef, fieldName string, tagName string, tagKey string) bool {
	tagVal := h.GetTagValByInputRef(inputRef, fieldName, tagName)
	if tagVal == "" {
		return false
	}

	// Check if the key even exists...
	if !strings.Contains(tagVal, tagKey) {
		return false
	}

	keys := strings.Split(tagVal, ";")
	if len(keys) == 0 {
		return false
	}
	// TODO: lowercase for checking...
	for _, v := range keys {
		keySplit := strings.Split(v, ":")
		keySplitLen := len(keySplit)
		if keySplitLen == 0 {
			continue
		}
		if keySplit[0] == tagKey {
			return true
		}
	}

	return false
}

func IsFieldTagKeyExists(obj interface{}, fieldName string, tagName string, tagKey string) bool {
	return New(obj).IsFieldTagKeyExists(nil, fieldName, tagName, tagKey)
}

func (h *Helper) GetFieldsByTagExistence(tagName string) []string {
	var fields []string

	for i := 0; i < h.valueOf.Type().NumField(); i++ {
		t := h.valueOf.Type().Field(i)
		fieldName := t.Name
		_, isFound := t.Tag.Lookup(tagName)

		if isFound {
			fields = append(fields, fieldName)
		}
	}
	return fields
}

func GetFieldsByTagExistence(obj interface{}, tagName string) []string {
	return New(obj).GetFieldsByTagExistence(tagName)
}

func (h *Helper) GetFieldNamesByTagKeyExistence(tagName string, tagKey string) []string {
	//if !IsPlainStruct(obj) {
	//	panic(InvalidObjType)
	//}

	var fields []string

	for i := 0; i < h.valueOf.Type().NumField(); i++ {
		t := h.valueOf.Type().Field(i)
		fieldName := t.Name
		tagVal := t.Tag.Get(tagName)

		if h.IsFieldTagKeyExists(&InputRef{TagVal: tagVal}, fieldName, tagName, tagKey) {
			fields = append(fields, fieldName)
		}
	}
	return fields
}

// GetFieldNamesByTagKeyExistence -> it will return the field names
func GetFieldNamesByTagKeyExistence(obj interface{}, tagName string, tagKey string) []string {
	return New(obj).GetFieldNamesByTagKeyExistence(tagName, tagKey)
}

/*
type FieldWithTagOptions struct {
	FieldName  string
	TagOptions []string
}

// GetFieldNamesByTagKeyExistenceWithAllTagOptions -> it will return the field names by indicated tag key + with all tag options
func GetFieldNamesByTagKeyExistenceWithAllTagOptions(obj interface{}, tagName string, tagKey string) []FieldWithTagOptions {
	val := reflect.ValueOf(obj)

	var fields []string

	for i := 0; i < val.Type().NumField(); i++ {
		t := val.Type().Field(i)
		fieldName := t.Name
		tagVal := t.Tag.Get(tagName)

		if IsFieldTagKeyExists(InputRef{TagVal: tagVal}, fieldName, tagName, tagKey) {
			fields = append(fields, fieldName)
		}
	}
	return fields
}*/

// tagOptions contains a slice of tag options
type tagOptions []string

// Has returns true if the given option is available in tagOptions
func (t tagOptions) Has(opt string) bool {
	for _, tagOpt := range t {
		if tagOpt == opt {
			return true
		}
	}

	return false
}

// parseTag splits a struct field's tag into its name and a list of options
// which comes after a name. A tag is in the form of: "name,option1,option2".
// The name can be neglectected.
func parseTag(tag string) (string, tagOptions) {
	// tag is one of followings:
	// ""
	// "name"
	// "name,opt"
	// "name,opt,opt2"
	// ",opt"

	res := strings.Split(tag, ",")
	return res[0], res[1:]
}
