package _struct

// IsStruct -> it can be a pointer of a struct or a struct!
// if it's a &Struct then it will return a ptr!, it's gonna be ptr
func IsStruct(obj interface{}) bool {
	return New(obj).isStr
}

func (h *Helper) isStruct() bool {
	if h.NonPtrObj == nil {
		return false
	}

	if h.typeOf == nil {
		return false
	}

	kind := h.typeOf.Kind().String()
	//log.Println("kind -> ", kind)
	if kind != "struct" {
		return false
	}
	return true
}

// IsPlainStruct -> checks if it's not a pointer and it's a clear/clean struct!
func (h *Helper) isPlainStruct() bool {
	if h.NonPtrObj == nil {
		return false
	}
	if h.isPtr || !h.isStr {
		return false
	}
	return true
}

// IsPlainStruct -> checks if it's not a pointer and it's a clear/clean struct!
func IsPlainStruct(obj interface{}) bool {
	return New(obj).isPlainStr
}
