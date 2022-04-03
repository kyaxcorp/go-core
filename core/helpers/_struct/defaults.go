package _struct

import "github.com/kyaxcorp/go-core/core/helpers/_struct/defaults"

// SetDefaultValues -> Set initializes members in a struct referenced by a pointer.
// Maps and slices are initialized by `make` and other primitive types are set with default values.
// `ptr` should be a struct pointer
func SetDefaultValues(obj interface{}) error {
	return defaults.Set(obj)
}
