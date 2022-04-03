package _struct

import "reflect"

const (
	DefaultTagName = "structs"
)

// structFields returns the exported struct fields for a given s struct. This
// is a convenient helper method to avoid duplicate code in some of the
// functions.
func (h *Helper) structFields() []reflect.StructField {
	t := h.valueOf.Type()

	var f []reflect.StructField

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// we can't access the value of unexported fields
		if field.PkgPath != "" {
			continue
		}

		// don't check if it's omitted
		if tag := field.Tag.Get(h.TagName); tag == "-" {
			continue
		}

		f = append(f, field)
	}

	return f
}
