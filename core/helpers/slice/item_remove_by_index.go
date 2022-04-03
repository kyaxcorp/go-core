package slice

/*
Remove by item
Remove by index

Remove and remain same indexes...
Remove and change indexes
*/

func RemoveByIndexInt(slice []int, index int, preserveOrder bool) []int {
	if preserveOrder {
		// it will preserve the order
		// But it will be much slower than the next method
		return append(slice[:index], slice[index+1:]...)
	} else {
		// This is the fastest method, but it doesn't preserve order
		slice[index] = slice[len(slice)-1]
		return slice[:len(slice)-1]
	}
}

func RemoveByIndexInt64(slice []int64, index int, preserveOrder bool) []int64 {
	if preserveOrder {
		// it will preserve the order
		// But it will be much slower than the next method
		return append(slice[:index], slice[index+1:]...)
	} else {
		// This is the fastest method, but it doesn't preserve order
		slice[index] = slice[len(slice)-1]
		return slice[:len(slice)-1]
	}
}

func RemoveByIndexInt32(slice []int32, index int, preserveOrder bool) []int32 {
	if preserveOrder {
		// it will preserve the order
		// But it will be much slower than the next method
		return append(slice[:index], slice[index+1:]...)
	} else {
		// This is the fastest method, but it doesn't preserve order
		slice[index] = slice[len(slice)-1]
		return slice[:len(slice)-1]
	}
}

func RemoveByIndexInt16(slice []int16, index int, preserveOrder bool) []int16 {
	if preserveOrder {
		// it will preserve the order
		// But it will be much slower than the next method
		return append(slice[:index], slice[index+1:]...)
	} else {
		// This is the fastest method, but it doesn't preserve order
		slice[index] = slice[len(slice)-1]
		return slice[:len(slice)-1]
	}
}

func RemoveByIndexInt8(slice []int8, index int, preserveOrder bool) []int8 {
	if preserveOrder {
		// it will preserve the order
		// But it will be much slower than the next method
		return append(slice[:index], slice[index+1:]...)
	} else {
		// This is the fastest method, but it doesn't preserve order
		slice[index] = slice[len(slice)-1]
		return slice[:len(slice)-1]
	}
}

func RemoveByIndexString(slice []string, index int, preserveOrder bool) []string {
	if preserveOrder {
		// it will preserve the order
		// But it will be much slower than the next method
		return append(slice[:index], slice[index+1:]...)
	} else {
		// This is the fastest method, but it doesn't preserve order
		slice[index] = slice[len(slice)-1]
		return slice[:len(slice)-1]
	}
}

func RemoveByIndexUint64(slice []uint64, index int, preserveOrder bool) []uint64 {
	if preserveOrder {
		// it will preserve the order
		// But it will be much slower than the next method
		return append(slice[:index], slice[index+1:]...)
	} else {
		// This is the fastest method, but it doesn't preserve order
		slice[index] = slice[len(slice)-1]
		return slice[:len(slice)-1]
	}
}

func RemoveByIndexUint(slice []uint, index int, preserveOrder bool) []uint {
	if preserveOrder {
		// it will preserve the order
		// But it will be much slower than the next method
		return append(slice[:index], slice[index+1:]...)
	} else {
		// This is the fastest method, but it doesn't preserve order
		slice[index] = slice[len(slice)-1]
		return slice[:len(slice)-1]
	}
}

func RemoveByIndexUint32(slice []uint32, index int, preserveOrder bool) []uint32 {
	if preserveOrder {
		// it will preserve the order
		// But it will be much slower than the next method
		return append(slice[:index], slice[index+1:]...)
	} else {
		// This is the fastest method, but it doesn't preserve order
		slice[index] = slice[len(slice)-1]
		return slice[:len(slice)-1]
	}
}

func RemoveByIndexUint16(slice []uint16, index int, preserveOrder bool) []uint16 {
	if preserveOrder {
		// it will preserve the order
		// But it will be much slower than the next method
		return append(slice[:index], slice[index+1:]...)
	} else {
		// This is the fastest method, but it doesn't preserve order
		slice[index] = slice[len(slice)-1]
		return slice[:len(slice)-1]
	}
}

func RemoveByIndexUint8(slice []uint8, index int, preserveOrder bool) []uint8 {
	if preserveOrder {
		// it will preserve the order
		// But it will be much slower than the next method
		return append(slice[:index], slice[index+1:]...)
	} else {
		// This is the fastest method, but it doesn't preserve order
		slice[index] = slice[len(slice)-1]
		return slice[:len(slice)-1]
	}
}

func RemoveByIndexBool(slice []bool, index int, preserveOrder bool) []bool {
	if preserveOrder {
		// it will preserve the order
		// But it will be much slower than the next method
		return append(slice[:index], slice[index+1:]...)
	} else {
		// This is the fastest method, but it doesn't preserve order
		slice[index] = slice[len(slice)-1]
		return slice[:len(slice)-1]
	}
}

func RemoveByIndexFloat64(slice []float64, index int, preserveOrder bool) []float64 {
	if preserveOrder {
		// it will preserve the order
		// But it will be much slower than the next method
		return append(slice[:index], slice[index+1:]...)
	} else {
		// This is the fastest method, but it doesn't preserve order
		slice[index] = slice[len(slice)-1]
		return slice[:len(slice)-1]
	}
}

func RemoveByIndexFloat32(slice []float32, index int, preserveOrder bool) []float32 {
	if preserveOrder {
		// it will preserve the order
		// But it will be much slower than the next method
		return append(slice[:index], slice[index+1:]...)
	} else {
		// This is the fastest method, but it doesn't preserve order
		slice[index] = slice[len(slice)-1]
		return slice[:len(slice)-1]
	}
}

func RemoveByIndexComplex128(slice []complex128, index int, preserveOrder bool) []complex128 {
	if preserveOrder {
		// it will preserve the order
		// But it will be much slower than the next method
		return append(slice[:index], slice[index+1:]...)
	} else {
		// This is the fastest method, but it doesn't preserve order
		slice[index] = slice[len(slice)-1]
		return slice[:len(slice)-1]
	}
}

func RemoveByIndexComplex64(slice []complex64, index int, preserveOrder bool) []complex64 {
	if preserveOrder {
		// it will preserve the order
		// But it will be much slower than the next method
		return append(slice[:index], slice[index+1:]...)
	} else {
		// This is the fastest method, but it doesn't preserve order
		slice[index] = slice[len(slice)-1]
		return slice[:len(slice)-1]
	}
}

func RemoveByIndexByte(slice []byte, index int, preserveOrder bool) []byte {
	if preserveOrder {
		// it will preserve the order
		// But it will be much slower than the next method
		return append(slice[:index], slice[index+1:]...)
	} else {
		// This is the fastest method, but it doesn't preserve order
		slice[index] = slice[len(slice)-1]
		return slice[:len(slice)-1]
	}
}
