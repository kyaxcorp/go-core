package Map

import "reflect"

func Clone(m map[interface{}]interface{}) map[interface{}]interface{} {
	nmap := make(map[interface{}]interface{})
	// TODO: well the cloning is not ideal here! because pointers may exist!
	for k, v := range m {
		nmap[k] = v
	}
	return nmap
}

func CloneStringInterface(m map[string]interface{}) map[string]interface{} {
	nmap := make(map[string]interface{})
	// TODO: well the cloning is not ideal here! because pointers may exist!
	for k, v := range m {
		nmap[k] = v
	}
	return nmap
}

func CopyStringInterface(from map[string]interface{}, to map[string]interface{}) {
	for k, v := range from {
		to[k] = v
	}
}

func IsMap(v interface{}) bool {
	if reflect.ValueOf(v).Kind() == reflect.Map {
		return true
	}
	return false
}
