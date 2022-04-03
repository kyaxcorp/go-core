package _struct

func StrStrToStrInterface(obj map[string]string) map[string]interface{} {
	if obj == nil {
		return nil
	}
	n := make(map[string]interface{})
	for k, v := range obj {
		n[k] = v
	}
	return n
}
