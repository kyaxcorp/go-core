package _struct

func (h *Helper) IsPointer() bool {
	return h.isPtr
}

func IsPointer(obj interface{}) bool {
	return New(obj).IsPointer()
}
